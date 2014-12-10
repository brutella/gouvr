package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestTime(t *testing.T) {
    tm := time.Now()
    ntm := NewTimeForUnixNano(time.Duration(tm.UnixNano()))
    assert.Equal(t, Milliseconds(ntm), Milliseconds(tm))
}