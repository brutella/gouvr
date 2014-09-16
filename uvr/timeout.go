package uvr

import (
    "time"
)

type Timeout struct {
    duration time.Duration
    deviation float64 // valid deviation (+/-) of a timeout between 0 and 1 (=100%) of the timeout duration
}

func NewTimeout(frequency time.Duration, deviation float64) Timeout {
    t := Timeout{}
    t.duration = time.Duration(time.Second/frequency)
    t.deviation = deviation
    
    return t
}

// Minimum valid timeout
func (t Timeout) min() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration - d
}

// Maximum valid timeout
func (t Timeout) max() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration + d
}

// Returns true when the timeout is already passed starting from a specified timestamp
func (t Timeout) IsPastSince(timestamp time.Time) bool {
    if time.Since(timestamp) > t.max() {
        return true
    }
    
    return false
}

// Returns ture when the timeout is already reached starting from a specified timestamp
func (t Timeout) IsFutureSince(timestamp time.Time) bool {
    if time.Since(timestamp) < t.min() {
        return true
    }
    
    return false
}