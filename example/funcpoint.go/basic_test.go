package funcpoint

import "testing"

func Test_BasicExample_1(t *testing.T) {
	expect := "Hello, world!"
	result := BasicExample()

	if expect != result {
		t.Fatalf("Failed! Test_BasicExample_1() out:%s", result)
	} else {
		t.Log("Pass! Test_BasicExample_1()")
	}
}
