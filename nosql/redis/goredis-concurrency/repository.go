package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"sync"
)

type ConcurrencyImplementation int

const (
	NoConcurrency ConcurrencyImplementation = iota
	AtomicOperator
	Transaction
	LUA
	Lock
)

type Repository interface {
	BuyShares(ctx context.Context, buyerId, companyId string, numShares int, wg *sync.WaitGroup) error
	GetCompanyShares(ctx context.Context, companyId string) (int, error)
	PublishShares(ctx context.Context, companyId string, numShares int) error
}

func BuildCompanySharesKey(companyId string) string {
	return fmt.Sprintf("gopp:shares:%s", companyId)
}

type SharesRepository struct {
	*redis.Client
	mutex *redsync.Mutex
}

func NewSharesRepository(address string) SharesRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	return SharesRepository{
		Client: client,
		mutex:  rs.NewMutex("test-redsync"),
	}
}

func (r *SharesRepository) BuyShares(ctx context.Context, impl ConcurrencyImplementation, userId, companyId string, numShares int, wg *sync.WaitGroup) error {
	defer wg.Done()
	switch impl {
	case NoConcurrency:
		return r.buySharesNoConcurrencyControl(ctx, userId, companyId, numShares)
	case AtomicOperator:
		return r.buySharesWithAtomicIncr(ctx, userId, companyId, numShares)
	case Transaction:
		return r.buySharesWithTransactions(ctx, userId, companyId, numShares)
	case LUA:
		return r.buySharesWithLUAScript(ctx, userId, companyId, numShares)
	case Lock:
		return r.buySharesWithRedisLock(ctx, userId, companyId, numShares)
	default:
		panic("invalid implementation method selectd")
	}
}

func (r *SharesRepository) buySharesNoConcurrencyControl(ctx context.Context, userId, companyId string, numShares int) error {
	var err error
	var iCurrentShares int
	// --- (1) ----
	// Get current number of shares
	currentShares := r.Get(ctx, BuildCompanySharesKey(companyId))
	if err = currentShares.Err(); err != nil {
		return err
	}

	// --- (2) ----
	// Validate if the shares remaining are enough to be bought
	iCurrentShares, err = currentShares.Int()
	if err != nil {
		return err
	}
	if iCurrentShares < numShares {
		return errors.New("error: company does not have enough shares")
	}
	iCurrentShares -= numShares

	// --- (3) ----
	// Update the current shares of the company and log who has bought shares
	if result := r.Set(ctx, BuildCompanySharesKey(companyId), iCurrentShares, 0); result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *SharesRepository) buySharesWithAtomicIncr(ctx context.Context, userId, companyId string, numShares int) error {
	var err error
	var iCurrentShares int
	// --- (1) ----
	// Get current number of shares
	currentShares := r.Get(ctx, BuildCompanySharesKey(companyId))
	if err = currentShares.Err(); err != nil {
		return err
	}

	// --- (2) ----
	// Validate if the shares remaining are enough to be bought
	iCurrentShares, err = currentShares.Int()
	if err != nil {
		return err
	}
	if iCurrentShares < numShares {
		return errors.New("error: company does not have enough shares")
	}

	// --- (3) ----
	// Update the current shares of the company and log who has bought shares
	if result := r.IncrBy(ctx, BuildCompanySharesKey(companyId), -1*int64(numShares)); result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *SharesRepository) buySharesWithTransactions(ctx context.Context, userId, companyId string, numShares int) error {
	var maxRetries = 100
	for i := 0; i < maxRetries; i++ {
		err := r.Watch(ctx, func(tx *redis.Tx) error {
			var err error
			var iCurrentShares int
			// --- (1) ----
			// Get current number of shares
			currentShares := r.Get(ctx, BuildCompanySharesKey(companyId))
			if err = currentShares.Err(); err != nil {
				return err
			}

			// --- (2) ----
			// Validate if the shares remaining are enough to be bought
			iCurrentShares, err = currentShares.Int()
			if err != nil {
				return err
			}
			if iCurrentShares < numShares {
				return errors.New("error: company does not have enough shares")
			}
			iCurrentShares -= numShares

			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				if result := pipe.Set(ctx, BuildCompanySharesKey(companyId), iCurrentShares, 0); result.Err() != nil {
					return result.Err()
				}
				return nil
			})

			return nil
		}, BuildCompanySharesKey(companyId))
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		return err
	}
	return errors.New("increment reached maximum number of retries")
}

var BuySharesLUA = redis.NewScript(`
local sharesKey = KEYS[1]
local requestedShares = ARGV[1]

local currentShares = redis.call("GET",sharesKey)
if currentShares < requestedShares then
	return {err = "error: company does not have enough shares"}
end

currentShares = currentShares - requestedShares
redis.call("SET",sharesKey,currentShares)
return currentShares
`)

func (r *SharesRepository) buySharesWithLUAScript(ctx context.Context, userId, companyId string, numShares int) error {
	keys := []string{BuildCompanySharesKey(companyId)}
	if result := BuySharesLUA.Run(ctx, r, keys, numShares); result.Err() != nil {
		if result.Err() == redis.Nil {
			return nil
		}
		return result.Err()
	}
	return nil
}

func (r *SharesRepository) buySharesWithRedisLock(ctx context.Context, userId, companyId string, numShares int) error {
	var err error
	if err = r.mutex.Lock(); err != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf("error during lock: %v \n", err)
		return err
	}
	defer func() {
		var ok bool
		if ok, err = r.mutex.Unlock(); !ok || err != nil {
			fmt.Println("error during unlock: %v \n", err)
			return
		}
	}()
	var iCurrentShares int
	// --- (1) ----
	// Get current number of shares
	currentShares := r.Get(ctx, BuildCompanySharesKey(companyId))
	if err := currentShares.Err(); err != nil {
		return err
	}

	// --- (2) ----
	// Validate if the shares remaining are enough to be bought
	iCurrentShares, err = currentShares.Int()
	if err != nil {
		return err
	}
	if iCurrentShares < numShares {
		return errors.New("error: company does not have enough shares")
	}
	iCurrentShares -= numShares
	// --- (3) ----
	// Update the current shares of the company and log who has bought shares
	if result := r.Set(ctx, BuildCompanySharesKey(companyId), iCurrentShares, 0); result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *SharesRepository) GetCompanyShares(ctx context.Context, companyId string) (int, error) {
	result := r.Get(ctx, BuildCompanySharesKey(companyId))
	currentShares, err := result.Int()
	if err != nil {
		return 0, err
	}
	return currentShares, nil
}

func (r *SharesRepository) PublishShares(ctx context.Context, companyId string, numShares int) error {
	status := r.Set(ctx, BuildCompanySharesKey(companyId), numShares, 0)
	return status.Err()
}
