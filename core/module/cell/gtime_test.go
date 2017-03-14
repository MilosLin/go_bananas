package cell

import (
	"testing"
	"time"

	"github.com/MilosLin/go_bananas/core/env"
)

func Test_Next5Min_1(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-08-05 00:00:00")
	if next5, err := Next5Min(currentTime); err != nil {
		t.Error("Test_Next5Min() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-08-05 00:05:00")
		if !expectTime.Equal(next5) {
			t.Error("Test_Next5Min() Failed")
		} else {
			t.Log("Test_Next5Min() Pass")
		}
	}
}

func Test_Next5Min_2(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-08-05 00:02:26")
	if next5, err := Next5Min(currentTime); err != nil {
		t.Error("Test_Next5Min() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-08-05 00:05:00")
		if !expectTime.Equal(next5) {
			t.Error("Test_Next5Min() Failed")
		} else {
			t.Log("Test_Next5Min() Pass")
		}
	}
}

func Test_Next5Min_3(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-08-05 23:59:59")
	if next5, err := Next5Min(currentTime); err != nil {
		t.Error("Test_Next5Min() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-08-06 00:00:00")
		if !expectTime.Equal(next5) {
			t.Error("Test_Next5Min() Failed")
		} else {
			t.Log("Test_Next5Min() Pass")
		}
	}
}

func Test_Next5Min_4(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-08-05 25:59:59")
	if _, err := Next5Min(currentTime); err != nil {
		t.Error("Test_Next5Min() Failed")
	} else {
		t.Log("Test_Next5Min() Pass")
	}
}
