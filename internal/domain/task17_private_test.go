package domain

import (
	"testing"
)

func TestTask17AuditNeedsActor(t *testing.T) {
	e := AuditEvent{Action: "state_changed", EntityID: "run-1"}
	if e.Task17Attributable() {
		t.Fatal("audit event without actor was accepted")
	}
}
