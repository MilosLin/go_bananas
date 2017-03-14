package cell

import (
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
