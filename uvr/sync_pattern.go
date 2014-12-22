package uvr

import (
	"math/big"
)

// SyncPattern contains info about expected bit value, timeouts, number of
// bits and is used to specify a sync pattern.
type SyncPattern struct {
	I       int
	Count   int
	Value   big.Word
	Timeout Timeout
	Last    *Bit
}
