package uvr

import (
    "fmt"
    "time"
)

type timeout struct {
    duration time.Duration
    deviation float64 // 0.0 - 1.0
}

func NewTimeout(frequency time.Duration, deviation float64) timeout {
    t := timeout{}
    t.duration = time.Duration(time.Second/frequency)
    t.deviation = deviation
    
    return t
}

func (t timeout) min() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration - d
}

func (t timeout) max() time.Duration {
    d := time.Duration(float64(t.duration) * t.deviation)
    return t.duration + d
}

func (t timeout) IsPastSince(timestamp time.Time) bool {
    if time.Since(timestamp) > t.max() {
        return true
    }
    
    return false
}

func (t timeout) IsFutureSince(timestamp time.Time) bool {
    if time.Since(timestamp) < t.min() {
        return true
    }
    
    return false
}

func (t timeout) PlausibleSince(timestamp time.Time) bool {
    duration := time.Since(timestamp)
    fmt.Sprintf("%d > %d && %d < %d", duration, t.min(), duration, t.max())
    if t.IsFutureSince(timestamp) == false && t.IsPastSince(timestamp) == false {
        return true
    }
    
    return false
}