package tool

import "strconv"

func StringToInt(i string) int {
	cnt, e := strconv.Atoi(i)
	if e != nil {
		cnt = 0
	}
	return cnt
}

func StringToUint32(s string) uint32 {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(u64)
}
