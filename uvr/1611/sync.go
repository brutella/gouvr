package uvr1611

import (
	_ "fmt"
	"github.com/brutella/gouvr/uvr"
	"math/big"
	"time"
)

type syncDecoder struct {
	uvr.BitConsumer
	bitConsumer  uvr.BitConsumer
	syncObserver uvr.SyncObserver
	synced       bool
	pattern      uvr.SyncPattern
}

// NewSyncDecoder returns a sync decoder, which implements the BitConsumer interface.
// After successful sync, the SyncDone() of the observer is called and the consumed
// bits are passed through to the specified bit consumer.
func NewSyncDecoder(bitConsumer uvr.BitConsumer, syncObserver uvr.SyncObserver, t uvr.Timeout) *syncDecoder {
	d := &syncDecoder{bitConsumer: bitConsumer, syncObserver: syncObserver}

	d.pattern = uvr.SyncPattern{
		Count:   32,
		Timeout: t,
	}

	return d
}

// Resets the cached bits and calls Reset() on the bitConsumer.
// After reset the decoder syncs the transmission again.
func (s *syncDecoder) Reset() {
	s.bitConsumer.Reset()
	s.resetBits()
}

func (s *syncDecoder) resetBits() {
	s.pattern.Last = nil
	s.pattern.I = 0 // reset
	s.synced = false
}

func (s *syncDecoder) Consume(bit uvr.Bit) error {
	if s.synced == true {
		// bitConsumer returns error when bit order is wrong
		// e.g. wrong start/stop bit
		err := s.bitConsumer.Consume(bit)
		if err != nil {
			s.Reset()
		}
	} else {
		// Check if the bit is within the allowed timeout
		// If bit arrived too early (should actually never happen), ignore it.
		// If bit arrived too late, return error.
		pattern := s.pattern
		if pattern.Last != nil {
			delta := time.Duration(bit.Timestamp.UnixNano() - pattern.Last.Timestamp.UnixNano())
			switch bit.CompareTimeoutToLast(pattern.Timeout, *pattern.Last) {
			case uvr.OrderedAscending:
				return nil
			case uvr.OrderedDescending:
				s.Reset()
				return uvr.NewErrorf("[SYNC] Bit arrived too late (%v)", delta)
			case uvr.OrderedSame:
			}
		}

		// Only accept low-high bit pattern
		if (pattern.Last == nil && bit.Raw == big.Word(0)) || (pattern.Last != nil && pattern.Last.Raw != bit.Raw) {
			s.pattern.I++
			s.pattern.Last = &bit
			if s.pattern.I == s.pattern.Count {
				if s.syncObserver != nil {
					s.syncObserver.SyncDone(bit.Timestamp)
				}
				s.synced = true
			}
		} else {
			s.Reset()
		}
	}
	return nil
}
