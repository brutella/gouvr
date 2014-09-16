package uvr

import (
    "time"
    "math/big"
    "strconv"
    "strings"
    "errors"
)

// A bit encapsulates a value and timestamp
type Bit struct {
    Raw big.Word // 1 or 0
    Timestamp time.Time // nanoseconds
}

// Creates a bit with a specific value and current timestamp
func NewBitFromWord(value big.Word) Bit {
    return Bit{Raw: value, Timestamp: time.Now()}
}

// Returns the time difference between bits
func (b Bit) Since(bit Bit) time.Duration {
    return time.Duration(b.Timestamp.UnixNano() - bit.Timestamp.UnixNano())
}

type ComparisonResult int
const(
    OrderedAscending    = iota // a < b
    OrderedSame         = iota // a = b
    OrderedDescending   = iota // a > b
)

// Compares the time difference between bits based on a timeout
func (b Bit) CompareTimeoutToLast(timeout Timeout, last Bit) ComparisonResult {
    elapsed := b.Since(last)
    if timeout.min() > elapsed {
        return OrderedAscending
    } else if timeout.max() < elapsed {
        return OrderedDescending
    }
    
    return OrderedSame
}

// Create a bit from a log string
// 
// The timestamp might be off by some nanoseconds due to problem converting string to int64.
func BitFromLogString(str string) (Bit, error) {
    var bit Bit
    var err error
    fields := strings.Fields(str)
    switch len(fields) {
    case 2:
        timestamp_str := fields[0]
        raw_str := fields[1]
        
        raw, _ := strconv.Atoi(raw_str)
        timestamp_nano, _ := strconv.ParseInt(timestamp_str, 10, 64)
        bit = Bit{Raw: big.Word(raw), Timestamp: NewTimeForUnixNano(time.Duration(timestamp_nano))}
    default:
        err = errors.New("Could not create bit from " + str)
    }
    
    return bit, err
}

// Creates a log string for a bit in the following format-
// 
// {nanoseconds} {1|0}
// ...
func LogString(b Bit) string {
    nanoSeconds := strconv.FormatInt(b.Timestamp.UnixNano(), 10)
    switch b.Raw {
    case 1:
        return nanoSeconds + " 1"
    default:
        return nanoSeconds + " 0"
    }
}