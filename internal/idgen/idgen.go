package idgen

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync/atomic"
	"time"
)

type Generator interface {
	New(prefix string) string
}

type Random struct{}

func (Random) New(prefix string) string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(fmt.Sprintf("generate random identifier: %v", err))
	}
	return prefix + "_" + hex.EncodeToString(b[:])
}

type Sequence struct {
	value atomic.Uint64
}

func (s *Sequence) New(prefix string) string {
	value := s.value.Add(1)
	return fmt.Sprintf("%s_%d_%06d", prefix, time.Now().UTC().Unix(), value)
}
