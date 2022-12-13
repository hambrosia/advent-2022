package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("\n%v \n is not equal to\n%v", a, b)
	}
}

func DebugPrint(enable bool, msg string, vars ...interface{}) {
	if enable {
		msg += "\n"
		fmt.Printf(msg, vars...)
	}
}
