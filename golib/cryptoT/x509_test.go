package cryptoT

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"
)

func Test_genPemFile(t *testing.T) {
	genPemFile("default.pem", "123456", []byte("654321"))
}

func Test_decodePemFile(t *testing.T) {
	log.Println(decodePemFile("default.pem", "123456"))
}

func genPemFile(filename, passwd string, data []byte) error {
	pemData, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", data, []byte(passwd), x509.PEMCipher3DES)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, pemData)
	if err != nil {
		return err
	}
	return nil
}

func decodePemFile(filename, passwd string) (pubkey, pirv interface{}, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	block, _ := pem.Decode(data)
	var pemData []byte
	pemData, err = x509.DecryptPEMBlock(block, []byte(passwd))
	pirv, err = x509.ParsePKCS1PrivateKey(pemData)
	if err != nil {
		return
	}
	pubkey = (pirv.(*rsa.PrivateKey)).PublicKey
	return
}

func Test(t *testing.T) {
	var list []Experiences
	list = append(list, newList([]int{1, 100, 0}, []int{2, 90, 0}, []int{3, 90, 100}, []int{4, 80, 100}))
	list = append(list, newList([]int{1, 100, 0}, []int{4, 80, 100}, []int{3, 90, 100}, []int{2, 90, 0}))
	list = append(list, newList([]int{4, 80, 100}, []int{3, 90, 100}, []int{2, 90, 0}, []int{1, 100, 0}))
	list = append(list, newList([]int{4, 80, 100}, []int{3, 90, 100}, []int{2, 90, 0}, []int{1, 100, 0}, []int{5, 60, 100}, []int{6, 60, 0}))
	for i := range list {
		log.Println("sort before->")
		printResult(list[i])
		sort.Stable(list[i])
		log.Println("sort after->")
		printResult(list[i])

		log.Println("end->")
	}
}

func printResult(list Experiences) {
	for i := range list {
		log.Println(fmt.Sprintf("Id:%d BeginTime:%v EndTime:%v", list[i].Id, list[i].BeginTime, list[i].EndTime))
	}
}

func newList(list ...[]int) Experiences {
	var rsp Experiences
	for i := range list {
		if len(list[i]) < 3 {
			continue
		}
		rsp = append(rsp, &Experience{Id: int(list[i][0]), BeginTime: int64(list[i][1]), EndTime: int64(list[i][2])})
	}
	return rsp
}

type Experiences []*Experience

func (t Experiences) Len() int {
	return len(t)
}

func (t Experiences) Less(i, j int) bool {
	if t[i].EndTime == 0 && t[j].EndTime == 0 {
		return t[i].BeginTime > t[j].BeginTime
	} else if t[i].EndTime == 0 && t[j].EndTime != 0 {
		return true
	} else if t[i].EndTime != 0 && t[j].EndTime != 0 {
		if t[i].EndTime != t[j].EndTime {
			return t[i].EndTime > t[j].EndTime
		} else {
			return t[i].BeginTime > t[j].BeginTime
		}
	} else if t[i].EndTime != 0 && t[j].EndTime == 0 {
		return false
	}
	return true
}

func (t Experiences) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type Experience struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`         //项目名称
	BeginTime    int64  `json:"begin_time"`   //开始时间
	EndTime      int64  `json:"end_time"`     //结束时间  0表示至今
	UpdateTime   int64  `json:"update_time"`  //更新时间
	Role         string `json:"role"`         //角色
	Introduction string `json:"introduction"` //项目描述
}
