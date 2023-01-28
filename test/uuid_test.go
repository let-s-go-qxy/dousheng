package test

import (
	"fmt"
	"testing"
	"tiktok/utils/uuid"
)

func TestUuid(t *testing.T) {
	uid := uuid.GetUUid()
	fmt.Println("uid:", uid, "\nlen:", len(uid))
	if len(uid) != 36 {
		t.Fail()
	}
}
