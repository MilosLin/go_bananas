package cell

import (
	"reflect"
	"testing"
)

func Test_GetIPv4(t *testing.T) {
	if reflect.TypeOf(GetIPv4()).String() != "net.IP" {
		t.Error("Test_GetIPv4() Failed")
	} else {
		t.Log("Test_GetIPv4() Pass")
	}
}
