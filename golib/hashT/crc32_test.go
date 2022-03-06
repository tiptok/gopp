package hashT

import (
	"crypto/md5"
	"hash/crc32"
	"log"
	"testing"
)

func TestCrc32(t *testing.T) {
	icrc := crc32.ChecksumIEEE([]byte("hello 2018"))
	log.Println("Crc32", icrc)

	hash32 := crc32.NewIEEE()
	hash32.Write([]byte("hello 2018"))
	log.Println("Crc32 sum32:", hash32.Sum32())
}

func BenchmarkCrc32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE([]byte("hello 2018"))
	}
}

//go test -bench.
func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5.Sum([]byte("hello 2018"))
	}
}
