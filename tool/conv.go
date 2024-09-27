package tool

import "strconv"

func StringToInt(i string) int {
	cnt, e := strconv.Atoi(i)
	if e != nil {
		cnt = 0
	}
	return cnt
}
