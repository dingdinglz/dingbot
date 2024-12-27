package database

import (
	"crypto/md5"
	"fmt"
	"io"
)

func md5Generate(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}
