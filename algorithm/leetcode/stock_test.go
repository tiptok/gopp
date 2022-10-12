package leetcode

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type DayStockInfo struct {
	Date   string
	Price  string
	InRate string
}

func newDayStockInfo(d ...string) *DayStockInfo {
	return &DayStockInfo{
		Date:   d[0],
		Price:  d[1],
		InRate: d[2],
	}
}

func monthDay() []*DayStockInfo {
	var stockInfo []*DayStockInfo
	stockInfo = append(stockInfo, newDayStockInfo("20190519", "6.12", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190520", "5.96", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190521", "5.95", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190522", "5.87", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190523", "5.80", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190524", "5.58", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190525", "5.60", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190526", "5.65", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190527", "5.72", "0.54"))
	stockInfo = append(stockInfo, newDayStockInfo("20190528", "5.67", "0.35"))
	return stockInfo
}

func Test_Stock(t *testing.T) {
	var stocks []*DayStockInfo = monthDay()
	AverangeModule(stocks)
}
func AverangeModule(stocks []*DayStockInfo) {
	log.Println("\nStock Average->")
	average := func(s []*DayStockInfo, peer int) {
		log.Println(fmt.Sprintf("Average Model: Peer %v Day->>", peer))
		totalSum := 0.0
		for i := 0; i < len(s); {
			sum := 0.0
			if i+peer > len(s) {
				break
			}
			begin := s[i]
			j := i
			var (
				lowest  float64
				highest float64
			)
			for j < len(s) && j < i+peer {
				tv, _ := strconv.ParseFloat(s[j].Price, 64)
				sum += tv
				j++
				if lowest == 0 && highest == 0 {
					lowest = tv
					highest = tv
				}
				if lowest > tv {
					lowest = tv
				}
				if highest < tv {
					highest = tv
				}
			}
			end := s[j-1]
			av := sum / float64(peer*1.0)
			totalSum += av
			log.Println(fmt.Sprintf("开始:%v 结束:%v 均值:%2.2f 最低:%2.2f 最高%2.2f", begin.Date, end.Date, av, lowest, highest))
			i += peer
		}
		log.Println(fmt.Sprintf("->Stock Average :%2.2f", totalSum/float64(len(stocks)/peer)))
	}

	//均值
	average(stocks, 3)
	average(stocks, 5)
	average(stocks, 10)

	//峰值
}
