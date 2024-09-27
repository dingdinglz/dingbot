package tool

import "os"

func FileIsExsits(_path string) bool {
	_, e := os.Stat(_path)
	return e == nil
}

func DictoryCreateN(_path string) {
	if !FileIsExsits(_path) {
		os.Mkdir(_path, os.ModePerm)
	}
}
