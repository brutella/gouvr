package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
	"testing"
)

func TestOutlets(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0x68), // 0110 1000
		uvr.Byte(0x00), // 0000 0000
	})
	outlets := OutletsFromValue(value)
	if is, want := len(outlets), 13; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[0].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[1].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[2].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[3].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[4].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[5].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[6].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[7].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[8].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[9].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[10].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[11].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[12].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
