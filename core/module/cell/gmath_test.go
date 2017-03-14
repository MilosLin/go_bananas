package cell

import "testing"

func Test_Abs(t *testing.T) {
	if (Abs(-5) != 5) || (Abs(0) != 0) || (Abs(9) != 9) {
		t.Error("Test_Abs() Failed")
	} else {
		t.Log("Test_Abs() Pass")
	}
}
