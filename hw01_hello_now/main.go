package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currentTime := time.Now()
	loc, _ := time.LoadLocation("UTC")
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal("Can not receive time with error: ", err)
	}

	fmt.Println("current time:", currentTime.In(loc))
	fmt.Println("exact time:", time.Local().In(loc))
}
