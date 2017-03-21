package funcpoint

import "testing"

func Test_CallBackExample_1(t *testing.T) {
	expect := "calculate(Plus):5\n"
	expect += "calculate(Minus):1\n"
	expect += "calculate(Multiply):6\n"
	result := CallBackExample()

	if expect != result {
		t.Fatalf("Failed! Test_CallBackExample_1() out:%s", result)
	} else {
		t.Log("Pass! Test_CallBackExample_1()")
	}
}
