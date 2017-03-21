package funcpoint

import "testing"

func Test_FlowControlExample_1(t *testing.T) {
	flow := 1
	expect := "A():inputB():input"
	result := FlowControlExample(flow)

	if expect != result {
		t.Fatalf("Failed! Test_FlowControlExample_1() out:%s", result)
	} else {
		t.Log("Pass! Test_FlowControlExample_1()")
	}
}

func Test_FlowControlExample_2(t *testing.T) {
	flow := 2
	expect := "A():inputC():input"
	result := FlowControlExample(flow)

	if expect != result {
		t.Fatalf("Failed! Test_FlowControlExample_2() out:%s", result)
	} else {
		t.Log("Pass! Test_FlowControlExample_2()")
	}
}

func Test_FlowControlExample_3(t *testing.T) {
	flow := 3
	expect := "A():inputB():inputC():input"
	result := FlowControlExample(flow)

	if expect != result {
		t.Fatalf("Failed! Test_FlowControlExample_3() out:%s", result)
	} else {
		t.Log("Pass! Test_FlowControlExample_3()")
	}
}
