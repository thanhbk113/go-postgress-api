package util

import "strconv"

func StringToInt(str string) int32 {
	num, _ := strconv.Atoi(str)
	return int32(num)
}
