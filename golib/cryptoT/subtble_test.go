package cryptoT

import (
	"crypto/subtle"
	"log"
	"testing"
)

func Test_subtle(t *testing.T) {
	log.Println(subtle.ConstantTimeEq(11, 11)) //equal 1  ,   other 0
}
