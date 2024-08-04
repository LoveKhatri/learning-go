package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	//                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^ -> The values need to be the same, format can be different
	fmt.Println(presentTime.Format("2006/02/01 04:15:05 Mon"))

	createdDate := time.Date(2026, time.July, 01, 12, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
