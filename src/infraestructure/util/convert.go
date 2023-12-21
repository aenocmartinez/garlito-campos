package util

import "strconv"

func ConverToInt64(value string) int64 {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return int64(number)
}
