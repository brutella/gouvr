package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
	"testing"
)

func TestHeatmetersEnabled(t *testing.T) {
	b := uvr.Byte(0x2) // 0000 0010
	h1, h2 := AreHeatMetersEnabled(uvr.NewHeatMeterRegister(b))

	if is, want := h1, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := h2, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
