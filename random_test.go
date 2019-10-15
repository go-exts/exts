package tool

import (
	"bytes"
	"testing"
)

func TestGenerateRandom(t *testing.T) {
	str := GenerateRandom(1024)
	for i := 0; i < 100000; i++ {
		str1 := GenerateRandom(1024)
		if bytes.Compare(str, str1) == 0 {
			str = str1
			t.Errorf("same:%s,%s", str, str1)
		}
	}
}
