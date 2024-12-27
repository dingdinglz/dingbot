package appconfig

import (
	"fmt"
	"testing"
)

func TestGenerateRandString(t *testing.T) {
	fmt.Println(GenerateRandString(10))
}
