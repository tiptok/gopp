package timeT

import (
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tNow := time.Now()
	if !tNow.IsZero() {
		log.Println("no zero time")
	}
	tZero, _ := time.Parse("2006-01-02 03:04:05", "0001-01-01 00:00:00")

	log.Println(tZero.Format("060102150405"))
	if tZero.IsZero() {
		log.Printf("%v is zero time", tZero)
	}
	log.Printf("%v %v is after:%v", tNow, tZero, tNow.After(tZero))
	log.Printf("%v %v is before:%v", tNow, tZero, tNow.Before(tZero))

	/*time.Duration 单位 ns*/
	tDurations := time.Duration(time.Second * 10)
	log.Printf("duration :%v %d", tDurations, tDurations)
	tDurations = time.Duration(time.Millisecond * 10)
	log.Printf("duration :%v %d", tDurations, tDurations)
}
