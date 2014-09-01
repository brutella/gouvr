package uvr

import (
    "time"
	"testing"
    "github.com/kidoman/embd"
    "github.com/stretchr/testify/assert"
)

type fakeDigitalPin struct {
    n int
}

func (p fakeDigitalPin) N() int  {
    return p.n
}

func (p fakeDigitalPin) Write(val int) error {
    return nil
}

func (p fakeDigitalPin) Read() (int, error) {
    return 1, nil
}

func (p fakeDigitalPin) TimePulse(state int) (time.Duration, error) {
    return 0, nil
}

func (p fakeDigitalPin) SetDirection(dir embd.Direction) error {
    return nil
}

func (p fakeDigitalPin) ActiveLow(b bool) error {
    return nil
}

func (p fakeDigitalPin) PullUp() error {
    return nil
}

func (p fakeDigitalPin) PullDown() error {
    return nil
}

func (p fakeDigitalPin) Close() error {
    return nil
}

func NewFakeDigitalPin(number int) (embd.DigitalPin, error) {
    return fakeDigitalPin{n: number}, nil
}

func NewFakeUVR1611() (UVR, error) {
    pin, _ := NewFakeDigitalPin(10)
    
    return New(pin, DeviceTypeUVR1611)
}

func TestCreateUVR1611(t *testing.T) {
    _, err := NewFakeUVR1611()
    assert.Nil(t, err)
    
}