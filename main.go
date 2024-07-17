package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {

	d := time.Now()
	h := time.Time.Hour(d)
	m := time.Time.Minute(d)

	defaultTime := d.Format("15:04")

	t := flag.String("t", defaultTime, "Set reference time")

	addh := flag.Int("h", 0, "Hours to add")
	addm := flag.Int("m", 0, "Minutes to add")

	flag.Parse()

	parsedTime, err := time.Parse("15:04", *t)
	if err != nil {
		fmt.Println(err)
	}

	parsedTimeStr := strconv.Itoa(time.Time.Hour(parsedTime)) + ":" + strconv.Itoa(time.Time.Minute(parsedTime))
	if defaultTime != parsedTimeStr {
		h = time.Time.Hour(parsedTime)
		m = time.Time.Minute(parsedTime)
	}

	h += *addh
	m += *addm

	final := time.Date(0, 0, 0, h, m, 0, 0, time.UTC)
	fmt.Println(final.Format("15:04"))
}
