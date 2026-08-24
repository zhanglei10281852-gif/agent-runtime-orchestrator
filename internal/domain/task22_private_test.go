package domain

import (
	"testing"
)

func TestTask22PoolAcceptsWithinCapacity(t *testing.T) {
	c := ExecutionPool{State: ExecutionPoolAvailable, CapacityUnits: 100}
	if !c.Task22CapacityAvailable(40) {
		t.Fatal("pool rejected capacity that fits")
	}
}
