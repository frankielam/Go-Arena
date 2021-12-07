package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)

	location := time.FixedZone("UTC+8", 0)		// TODO
	then := time.Date(2021, 12, 8, 11, 11, 11, 100000000, location)	// year, month, day, hour, minute, second, nanosecond, TZ
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())


	p(then, "before", now, then.Before(now))
	p(then, "after", now, then.After(now))
	p(then, "equal", now, then.Equal(now))

	diff := then.Sub(now)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(now.Add(diff))
	p(now.Add(-diff))


	firstDay := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	p(firstDay)

	days := now.Sub(firstDay)
	p(days.Hours())
	p(days.Hours()/24)
	p(days.Hours()/24/7)

	p(firstDay.Weekday())
	p(now.Weekday())
}