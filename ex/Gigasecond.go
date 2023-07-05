package ex

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigaSecond := 1000000000

	return t.Add(time.Duration(gigaSecond) * time.Second)
}
	 	