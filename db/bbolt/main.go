package main

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

func main() {
	// 打开数据库文件，如果文件不存在则创建
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建一个桶（类似于关系数据库中的表）
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 插入键值对
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("key1"), []byte("value1"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// 获取键值对
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("key1"))
		fmt.Printf("Value of key1: %s\n", string(v))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 删除键值对
	//err = db.Update(func(tx *bbolt.Tx) error {
	//	b := tx.Bucket([]byte("MyBucket"))
	//	err := b.Delete([]byte("key1"))
	//	return err
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 再次获取已删除的键值对
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("key1"))
		fmt.Printf("Value of key1 after deletion: %v\n", v)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 查询桶中的所有值
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		if b == nil {
			return fmt.Errorf("桶不存在")
		}
		c := b.Cursor()

		fmt.Println("桶中的所有值:")
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("键: %s, 值: %s\n", string(k), string(v))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("查询数据失败: %v", err)
	}
}
