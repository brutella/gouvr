package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestTimeout(t *testing.T) {
    duration := time.Duration(2)
    deviation := 0.5 // 50%
    timeout := NewTimeout(duration, deviation)
    
    assert.Equal(t, timeout.min(), time.Duration(250 * time.Millisecond))
    assert.Equal(t, timeout.max(), time.Duration(750 * time.Millisecond))
    
    now := time.Now()
    assert.True(t, timeout.IsFutureSince(now))
    assert.False(t, timeout.IsPastSince(now))
    
    time.Sleep(250 * time.Millisecond)
    assert.False(t, timeout.IsFutureSince(now))
    assert.False(t, timeout.IsPastSince(now))
    
    time.Sleep(600 * time.Millisecond)
    assert.False(t, timeout.IsFutureSince(now))
    assert.True(t, timeout.IsPastSince(now))
}
