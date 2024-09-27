package tool

import (
	"os"
	"path/filepath"
	"time"
)

func ErrorOut(part string, message string, err error) {
	rootPath, _ := os.Getwd()
	var endOut string = "[" + part + "]" + time.Now().String() + "\n[Message]" + message + "\n" + err.Error()
	os.WriteFile(filepath.Join(rootPath, "ERROR.LOG"), []byte(endOut), os.ModePerm)
}
