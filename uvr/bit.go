package uvr

import (
    "time"
    "math/big"
    "strconv"
    "strings"
    "errors"
)

type Bit struct {
    Raw big.Word
    Timestamp time.Time
}

func NewBitFromWord(value big.Word) Bit {
    return Bit{Raw: value, Timestamp: time.Now()}
}

func NewTimeForUnixNano(unixNano time.Duration) time.Time {
    seconds_f := float64(unixNano)/float64(time.Second)
    seconds_int := float64(int64(seconds_f))
    nanoseconds_int := (seconds_f - seconds_int) * float64(time.Second)        
    return time.Unix(int64(seconds_int), int64(nanoseconds_int))
}

func (b Bit) Since(bit Bit) time.Duration {
    return time.Duration(b.Timestamp.UnixNano() - bit.Timestamp.UnixNano())
}

type ComparisonResult int
const(
    OrderedAscending    = iota
    OrderedSame         = iota
    OrderedDescending   = iota
)

func (b Bit) CompareTimeoutToLast(timeout Timeout, last Bit) ComparisonResult {
    elapsed := b.Since(last)
    if timeout.min() > elapsed {
        return OrderedAscending
    } else if timeout.max() < elapsed {
        return OrderedDescending
    }
    
    return OrderedSame
}

/** Create a bit from a log string
    
    The timestamp might be off by some nanoseconds
    due to problem converting string to int64.
*/
func BitFromLogString(str string) (Bit, error) {
    var bit Bit
    var err error
    fields := strings.Fields(str)
    switch len(fields) {
    case 2:
        timestamp_str := fields[0]
        raw_str := fields[1]
        
        timestamp, _ := strconv.ParseInt(timestamp_str, 10, 64)
        raw, _ := strconv.Atoi(raw_str)
        seconds_f := float64(timestamp)/float64(time.Second)
        seconds_int := float64(int64(seconds_f))
        nanoseconds_int := (seconds_f - seconds_int) * float64(time.Second)        
        bit = Bit{Raw: big.Word(raw), Timestamp: time.Unix(int64(seconds_int), int64(nanoseconds_int))}
    default:
        err = errors.New("Could not create bit from " + str)
    }
    
    return bit, err
}

/** Creates a log string for a bit in the following format-
    
    {nanoseconds} {1|0}
*/
func LogString(b Bit) string {
    nanoSeconds := strconv.FormatInt(b.Timestamp.UnixNano(), 10)
    switch b.Raw {
    case 1:
        return nanoSeconds + " 1"
    default:
        return nanoSeconds + " 0"
    }
}