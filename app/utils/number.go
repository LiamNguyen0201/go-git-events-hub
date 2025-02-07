package utils

import (
	"log"
	"strconv"
)

func StringToNumber(input string) int64 {
	// Convert string to int64
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println("(StringToNumber) Input: ", input)
		log.Println("(StringToNumber) Error: ", err)
		panic("(StringToNumber) Convert string to int64 failed!")
	}

	return num
}
