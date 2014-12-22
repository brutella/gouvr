package uvr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tm := time.Now()
	ntm := NewTimeForUnixNano(time.Duration(tm.UnixNano()))
	assert.Equal(t, Milliseconds(ntm), Milliseconds(tm))
}
