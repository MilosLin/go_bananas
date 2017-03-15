package cell

import (
	"fmt"
	"strconv"
	"time"
)

// Next5Min 取下一個5分鐘時間點
//
// 例如 input: 03:23 、 output: 05:00
//
//   |------|------|-----------|
// 00:00  03:23  05:00       10:00
func Next5Min(current_time time.Time) (time.Time, error) {
	add_min := Abs((current_time.Minute() % 5) - 5)
	min_str := "+" + strconv.Itoa(add_min) + "m"
	if min_dur, err := time.ParseDuration(min_str); err != nil {
		return current_time, err
	} else {
		current_time = current_time.Add(min_dur)
	}

	//秒數固定為00
	sec_str := "-" + strconv.Itoa(current_time.Second()) + "s"
	if deduct_sec, err := time.ParseDuration(sec_str); err != nil {
		return current_time, err
	} else {
		current_time = current_time.Add(deduct_sec)
	}
	return current_time, nil
}

// 取當周周一，開始時間
//
// ex. input: 2017-03-15 16:12:08 output:2017-03-13 00:00:00
func CurrentWeekStart(current_time time.Time) (time.Time, error) {
	passByDays := [7]int{6, 0, 1, 2, 3, 4, 5}
	passByDay := passByDays[int(current_time.Weekday())]

	backtrack_str := fmt.Sprintf(
		"-%dh%dm%ds",
		current_time.Hour()+passByDay*24,
		current_time.Minute(),
		current_time.Second(),
	)
	if backtrack_dur, err := time.ParseDuration(backtrack_str); err != nil {
		return current_time, err
	} else {
		current_time = current_time.Add(backtrack_dur)
	}

	return current_time, nil
}
