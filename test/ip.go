package main

import (
	"log"

	"github.com/yinheli/qqwry"
)

func main() {
	q := qqwry.NewQQwry("/Users/zoucaitou/Downloads/qqwry.dat")
	q.Find("119.98.193.162")
	log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
}
