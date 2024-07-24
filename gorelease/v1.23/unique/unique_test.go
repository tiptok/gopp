package unique

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"testing"
)

func wordGen(nDistinct, wordLen int) func() string {
	vocab := make([]string, nDistinct)
	for i := range nDistinct {
		word := randomString(wordLen)
		vocab[i] = word
	}
	return func() string {
		word := vocab[rand.Intn(nDistinct)]
		return strings.Clone(word)
	}
}

func randomString(n int) string {
	const letters = "eddycjyabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; {
		b := make([]byte, 1)
		if _, err := rand.Read(b); err != nil {
			panic(err)
		}
		ret[i] = letters[int(b[0])%len(letters)]
		i++
	}
	return string(ret)
}

const nWords = 10000
const nDistinct = 100
const wordLen = 40

var words []string

// Memory used: 460KB
func Test_Common_WordGen(t *testing.T) {
	words = make([]string, nWords)
	generate := wordGen(nDistinct, wordLen)
	method := func() {
		for i := range nWords {
			words[i] = generate()
		}
	}
	computeAlloc(method)
}

//var wordsUnique []unique.Handle[string]
//// Memory used: 95KB
//func Test_Unique_WordGen(t *testing.T) {
//	wordsUnique = make([]string, nWords)
//	generate := wordGen(nDistinct, wordLen)
//	method := func() {
//		for i := range nWords {
//			wordsUnique[i] = generate()
//		}
//	}
//	computeAlloc(method)
//}

func computeAlloc(method func()) {
	memBefore := getAlloc()
	method()
	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	if memBefore > memAfter {
		memUsed = memBefore - memAfter
	}
	fmt.Printf("Memory used: %dKB \n", memUsed/1024)
}

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
