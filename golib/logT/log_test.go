package logT

import (
	"log"
	"os"
	"testing"
)

func Test_Log(t *testing.T) {
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 666)
	defer file.Close()
	log.SetOutput(file)
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	log.SetPrefix("[Test] ")
	log.Println("on test 3")
	log.Println("on test 4")
}
