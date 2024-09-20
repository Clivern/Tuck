// Time
//
// Go offers extensive support for times and durations; here are some examples.
//
// Run: go run ./64_time
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// We'll start by getting the current time. Its value changes on every run, so we
	// only print its type here.
	now := time.Now()
	fmt.Printf("now is a %T\n", now)

	// You can build a time struct by providing the year, month, day, etc. Times are
	// always associated with a Location, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday Weekday is also available.
	p(then.Weekday())

	// These methods compare two times, testing if the first occurs before, after, or
	// at the same time as the second, respectively.
	later := then.Add(48 * time.Hour)
	p(then.Before(later))
	p(then.After(later))
	p(then.Equal(later))

	// The Sub methods returns a Duration representing the interval between two
	// times.
	diff := later.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with a - to move
	// backwards by a duration.
	p(then.Add(diff))
	p(then.Add(-diff))

	// AddDate moves by calendar units, and Truncate rounds down to a multiple of a
	// duration.
	p(then.AddDate(0, 1, 0))
	p(then.Truncate(time.Hour))

	// Durations parse from strings like "1h30m", and convert between zones with In.
	d, _ := time.ParseDuration("1h15m30.5s")
	p(d, d.Round(time.Minute))
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err == nil {
		p(then.In(tokyo))
	}

	// time.Since measures elapsed time; it's the usual way to time a piece of code.
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	p("elapsed at least 10ms:", time.Since(start) >= 10*time.Millisecond)
}
