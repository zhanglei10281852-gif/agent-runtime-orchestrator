package domain

import (
	"testing"
	"time"
)

func TestTask21ApprovalDeadlineClosesTask(t *testing.T) {
	now := time.Date(2026, 8, 24, 9, 0, 0, 0, time.UTC)
	h := ApprovalTask{Status: ApprovalTaskPending, ExpiresAt: now}
	if h.Task21CanResolveAt(now) {
		t.Fatal("approval task stayed open at deadline")
	}
}
