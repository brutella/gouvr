package uvr

import (
	"fmt"
	"math/big"
	"time"
)

type syncDecoder struct {
	BitConsumer
	bitConsumer  BitConsumer
	syncObserver SyncObserver
	synced       bool
	pattern      SyncPattern
}

// NewSyncDecoder returns a bit consumer which decodes a bit stream to synchronize the communication
// The sync pattern is 8 high bits after a specific timeout.
//
// The BitConsumer's Consume method is called after synchronization. If an error is returned, the
// BitConsumer's Reset method is called.
// The SyncObserver's SyncDone method is called after synchronization.
// The Timeout specifies the time between two bits.
func NewSyncDecoder(bitConsumer BitConsumer, syncObserver SyncObserver, t Timeout) *syncDecoder {
	d := &syncDecoder{bitConsumer: bitConsumer, syncObserver: syncObserver}

	d.pattern = SyncPattern{
		Count:   8,
		Value:   big.Word(1),
		Timeout: t,
	}

	return d
}

func (s *syncDecoder) resetBits() {
	s.pattern.Last = nil
	s.pattern.I = 0
	s.synced = false
}
func (s *syncDecoder) Reset() {
	s.bitConsumer.Reset()
	s.resetBits()
}

func (s *syncDecoder) Consume(bit Bit) error {
	if s.synced == true {
		// BitConsumer returns error when bit order is wrong
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
			case OrderedAscending:
				fmt.Printf("[SYNC] Bit arrived too early (%v)\n", delta)
				return nil
			case OrderedDescending:
				s.Reset()
				err := NewErrorf("[SYNC] Bit arrived too late (%v)", delta)
				return err
			case OrderedSame:
			}
		}
		// Only accept pattern bits (e.g. only high bits)
		if bit.Raw == s.pattern.Value {
			s.pattern.I++
			s.pattern.Last = &bit
			if s.pattern.I == s.pattern.Count {
				if s.syncObserver != nil {
					s.syncObserver.SyncDone(bit.Timestamp)
				}
				s.synced = true
				fmt.Println("[SYNC] Done")
			}
		} else {
			s.resetBits()
		}
	}
	return nil
}
