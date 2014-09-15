package uvr

import (
    "time"
)

type Timeout struct {
    duration time.Duration
    deviation float64 // 0.0 - 1.0
}

func NewTimeout(frequency time.Duration, deviation float64) Timeout {
    t := Timeout{}
    t.duration = time.Duration(time.Second/frequency)
    t.deviation = deviation
    
    return t
}

func (t Timeout) min() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration - d
}

func (t Timeout) max() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration + d
}

func (t Timeout) IsPastSince(timestamp time.Time) bool {
    if time.Since(timestamp) > t.max() {
        return true
    }
    
    return false
}

func (t Timeout) IsFutureSince(timestamp time.Time) bool {
    if time.Since(timestamp) < t.min() {
        return true
    }
    
    return false
}