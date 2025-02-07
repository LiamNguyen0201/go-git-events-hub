package utils

import (
	"strconv"
)

func StringToNumber(input string) int64 {
	// Convert string to int64
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		LogDebug("(StringToNumber) Input: ", input)
		LogDebug("(StringToNumber) Error: ", err)
		panic("(StringToNumber) Convert string to int64 failed!")
	}

	return num
}
