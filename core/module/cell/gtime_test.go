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

func Test_CurrentWeekStart_1(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-03-15 16:12:08")
	if monday, err := CurrentWeekStart(currentTime); err != nil {
		t.Error("Test_CurrentWeekStart_1() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-03-13 00:00:00")
		if !expectTime.Equal(monday) {
			t.Error("Test_CurrentWeekStart_1() Failed")
		} else {
			t.Log("Test_CurrentWeekStart_1() Pass")
		}
	}
}

func Test_CurrentWeekStart_2(t *testing.T) {
	taipei_location, _ := time.LoadLocation("Asia/Taipei")
	currentTime, _ := time.ParseInLocation(env.DateTime, "2017-03-15 16:12:08", taipei_location)

	t.Log(currentTime.String())

	if monday, err := CurrentWeekStart(currentTime); err != nil {
		t.Error("Test_CurrentWeekStart_2() Failed")
	} else {
		expectTime, _ := time.ParseInLocation(env.DateTime, "2017-03-13 00:00:00", taipei_location)
		if !expectTime.Equal(monday) {
			t.Error("Test_CurrentWeekStart_2() Failed")
		} else {
			t.Log("Test_CurrentWeekStart_2() Pass")
		}
	}
}

func Test_CurrentWeekStart_3(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-03-20 23:59:59")
	if monday, err := CurrentWeekStart(currentTime); err != nil {
		t.Error("Test_CurrentWeekStart_3() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-03-20 00:00:00")
		if !expectTime.Equal(monday) {
			t.Error("Test_CurrentWeekStart_3() Failed")
		} else {
			t.Log("Test_CurrentWeekStart_3() Pass")
		}
	}
}

func Test_CurrentWeekStart_4(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-03-20 00:00:00")
	if monday, err := CurrentWeekStart(currentTime); err != nil {
		t.Error("Test_CurrentWeekStart_4() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-03-20 00:00:00")
		if !expectTime.Equal(monday) {
			t.Error("Test_CurrentWeekStart_4() Failed")
		} else {
			t.Log("Test_CurrentWeekStart_4() Pass")
		}
	}
}

func Test_CurrentWeekStart_5(t *testing.T) {
	currentTime, _ := time.Parse(env.DateTime, "2017-03-19 23:59:59")
	if monday, err := CurrentWeekStart(currentTime); err != nil {
		t.Error("Test_CurrentWeekStart_5() Failed")
	} else {
		expectTime, _ := time.Parse(env.DateTime, "2017-03-13 00:00:00")
		if !expectTime.Equal(monday) {
			t.Error("Test_CurrentWeekStart_5() Failed")
		} else {
			t.Log("Test_CurrentWeekStart_5() Pass")
		}
	}
}
