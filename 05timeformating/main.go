package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time formating in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	creatTime := time.Date(2001, time.September, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(creatTime.Format("01-02-2006 15:04:05 Monday"))

}
