package cryptoT

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("tiptok"))
	h.Write([]byte("ccc"))
	log.Println(hex.EncodeToString(h.Sum(nil)))

	log.Println("Home Dri:" + UserHomeDir())
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		log.Println("HOMEDRIVE:" + os.Getenv("HOMEDRIVE") + " HOMEPATH:" + os.Getenv("HOMEPATH"))
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

//bcrypt 存储密码
func TestBcrypt(t *testing.T) {
	password := "123456"
	passwordERR := "12345678"
	md5Pass := genMd5Hex([]byte(password))
	log.Println("gen md5:", hex.EncodeToString(md5Pass))
	hashPass, err := bcrypt.GenerateFromPassword(md5Pass, bcrypt.DefaultCost) //保存 hashPass用与以后比较
	if err != nil {
		t.Fatal(err)
	}
	log.Println("bcrypt:", string(hashPass))
	err = bcrypt.CompareHashAndPassword(hashPass, md5Pass)
	if err == nil {
		log.Println("password correct.")
	}
	err = bcrypt.CompareHashAndPassword(hashPass, genMd5Hex([]byte(passwordERR)))
	if err != nil {
		log.Println("password err.", err)
	}
}

func genMd5Hex(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}
