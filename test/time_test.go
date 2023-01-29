package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now().UnixNano()
	fmt.Println(now)
	if str := reflect.TypeOf(now); str.String() != "int64" {
		fmt.Println(str)
		t.Fail()
	}
}
