package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
)

var selectedImpl = flag.Int("i", 4, "选择具体实现执行的程序")
var totalClient = flag.Int("c", 30, "并非数")
var totalNumberShares = flag.Int("s", 1000, "累计股票数量")

const companyId = "TestCompanySL"

func main() {
	flag.Parse()
	switch ConcurrencyImplementation(*selectedImpl) {
	case NoConcurrency:
		fmt.Println(">> No Concurrency selected...")
	case AtomicOperator:
		fmt.Println(">> Atomic Operator selected...")
	case Transaction:
		fmt.Println(">> Transaction selected...")
	case LUA:
		fmt.Println(">> LUA Script selected...")
	case Lock:
		fmt.Println(">> Redis Locks selected...")
	default:
		panic("invalid implementation method selected")
	}
	// 对于redis并发问题，需要采用LUA、Lock才可以解决，LUA性能更高，只是会阻塞其他事务

	// --- (1) ----
	// Get the redis config and init the repository
	repository := NewSharesRepository("0.0.0.0:6379")

	// --- (2) ----
	// Publish available shares
	repository.PublishShares(context.Background(), companyId, *totalNumberShares)

	// --- (3) ----
	// Run concurrent clients that buy shares
	var wg sync.WaitGroup
	wg.Add(*totalClient)

	for idx := 1; idx <= *totalClient; idx++ {
		userId := fmt.Sprintf("user%d", idx)
		go func() {
			if err := repository.BuyShares(context.Background(), ConcurrencyImplementation(*selectedImpl), userId, companyId, 100, &wg); err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
	wg.Wait()

	// --- (3) ----
	// Get the remaining company shares
	shares, err := repository.GetCompanyShares(context.Background(), companyId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the number of free shares the company %s has is: %d", companyId, shares)
}
