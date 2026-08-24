package domain

import (
	"testing"
)

func TestTask06PoolAcceptsWithinCapacity(t *testing.T) {
	c := ExecutionPool{State: ExecutionPoolAvailable, CapacityUnits: 100}
	if !c.Task06CapacityAvailable(40) {
		t.Fatal("pool rejected capacity that fits")
	}
}
