package uvr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {
	bytes := []Byte{Byte(0x11), Byte(0xFA)}
	value := NewValue(bytes)

	// 0xFA11 = 64017
	assert.Equal(t, UInt16FromValue(value), uint16(64017))
	// 0xFA11 = -31249 unsigned
	assert.Equal(t, Int16FromValue(value), -((65536 / 2) - 31249))
}
