package uvr

import (
    "time"
)

func NewTimeForUnixNano(unixNano time.Duration) time.Time {
    seconds_f := float64(unixNano)/float64(time.Second)
    seconds_int := float64(int64(seconds_f))
    nanoseconds_int := (seconds_f - seconds_int) * float64(time.Second)        
    return time.Unix(int64(seconds_int), int64(nanoseconds_int))
}
