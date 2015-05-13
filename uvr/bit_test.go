package uvr

import (
	"math/big"
	"testing"
	"time"
)

func milliseconds(t time.Time) int64 {
	return int64((t.UnixNano() % 1e9) / 1e6)
}

func TestLogStringFromBit(t *testing.T) {
	bit := NewBitFromWord(1)
	str := LogString(bit)
	lbit, err := BitFromLogString(str)

	if err != nil {
		t.Fatal(err)
	}
	if is, want := lbit.Raw, bit.Raw; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	// Only test milliseconds
	if is, want := milliseconds(lbit.Timestamp), milliseconds(bit.Timestamp); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestBitsTimeout(t *testing.T) {
	bit1 := Bit{Raw: big.Word(1), Timestamp: time.Unix(0, 1001)}
	bit0 := Bit{Raw: big.Word(0), Timestamp: time.Unix(0, 1002)}

	tout := Timeout{duration: time.Duration(1), deviation: 0}
	// time difference is 1002 - 1001 = 1
	// valid timeout = +/-1
	if is, want := bit0.CompareTimeoutToLast(tout, bit1), OrderedSame; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestBitsTimeoutDeviation(t *testing.T) {
	bit1 := Bit{Raw: big.Word(1), Timestamp: time.Unix(0, 1001)}
	bit0 := Bit{Raw: big.Word(0), Timestamp: time.Unix(0, 1003)}

	tout := Timeout{duration: time.Duration(1), deviation: 1}
	// time difference is 1003 - 1001 = 2
	// valid timeout = +/-2
	if is, want := bit0.CompareTimeoutToLast(tout, bit1), OrderedSame; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestBitsOutOfTimeout(t *testing.T) {
	bit1 := Bit{Raw: big.Word(1), Timestamp: time.Unix(0, 1001)}
	bit0 := Bit{Raw: big.Word(0), Timestamp: time.Unix(0, 1003)}

	tout := Timeout{duration: time.Duration(1), deviation: 0}
	// time difference is 1003 - 1001 = 2
	// valid timeout = +/-1
	if is, want := bit0.CompareTimeoutToLast(tout, bit1), OrderedDescending; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
