package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("123456")
	md5Ret := md5.Sum(data)
	has := fmt.Sprintf("%x", md5Ret)
	fmt.Println(has)
	if len(has) != 32 {
		t.Fail()
	}
}
