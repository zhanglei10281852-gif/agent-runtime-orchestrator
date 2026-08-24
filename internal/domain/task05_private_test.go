package domain

import (
	"testing"
)

func TestTask05PoolAcceptsWithinCapacity(t *testing.T) {
	c := ExecutionPool{State: ExecutionPoolAvailable, CapacityUnits: 100}
	if !c.Task05CapacityAvailable(40) {
		t.Fatal("pool rejected capacity that fits")
	}
}
