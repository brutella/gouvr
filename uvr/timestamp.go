package uvr

import (
    "fmt"
)

type Timestamp struct {
    daylightSavings bool
    minute  Byte
    hour    Byte
    day     Byte
    month   Byte
    year    Byte // since 2000, e.g. 3 == 2003
}

/** Documentation
    Die Sommerzeit wird durch das Bit 5 im Stunden-Byte gekennzeichnet (bei Sommerzeit ist dieses Bit = 1).
    Zu bemerken ist, dass die akt. Uhrzeit auf den niederwertigen 5 Bits ausgegeben wird.
*/
func NewTimestamp(bytes []Byte) Timestamp {
    daylightSavings := bytes[1] & 0x20  == 0x20 // 0010 0000
    hour := bytes[1] & 0x1F // 0001 1111
    return Timestamp{
        daylightSavings: daylightSavings,
        minute: bytes[0],
        hour: hour,
        day: bytes[2],
        month: bytes[3],
        year: bytes[4],
    }
}

func (t Timestamp) ToString() string {
    return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00", 2000 + int64(t.year), t.month, t.day, t.hour, t.minute)
}