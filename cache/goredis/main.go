package main

import (
	"fmt"
	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
	"time"
)

func main() {
	Example_basicUsage()
}

type Object struct {
	Str string
	Num int
}

func Example_basicUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"127.0.0.1": ":6379",
			//"server2": ":6380",
		},
		Password: "123456",
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output: {mystring 42}
}

func Example_advancedUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	obj := new(Object)
	err := codec.Once(&cache.Item{
		Key:    "mykey",
		Object: obj, // destination
		Func: func() (interface{}, error) {
			return &Object{
				Str: "mystring",
				Num: 42,
			}, nil
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	// Output: &{mystring 42}
}
