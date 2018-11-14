package main

import (
	"fmt"
	"time"
)

func zeroTime() time.Time {
	zero, err := time.Parse("January 2, year 2006, 15:04:05 MST", "January 1, year 1, 00:00:00 UTC")
	if err != nil {
		panic(err)
	}

	return zero
}

func isZero(checkValue time.Time) bool {
	return checkValue.IsZero()
}

func main() {
	fmt.Println(isZero(zeroTime()))
}
