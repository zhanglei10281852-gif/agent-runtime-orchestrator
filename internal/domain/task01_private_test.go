package domain

import (
	"testing"
	"time"
)

func TestTask01SessionRevocationBoundary(t *testing.T) {
	now := time.Date(2026, 8, 24, 9, 0, 0, 0, time.UTC)
	revokedAt := now.Add(-time.Minute)
	s := Session{ExpiresAt: now.Add(time.Hour), RevokedAt: &revokedAt}
	if s.Task01ActiveAt(now) {
		t.Fatal("revoked session remained active")
	}
}
