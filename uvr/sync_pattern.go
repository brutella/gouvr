package uvr

import(
    "math/big"
)

// Holds info about sync pattern like expected bit value, timeouts, number of bits,...
type SyncPattern struct {
    I int
    Count int
    Value big.Word
    Timeout Timeout
    Last *Bit
}