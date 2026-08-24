package domain

import (
	"testing"
	"time"
)

func TestTask13ExactWindowLimitIsAccepted(t *testing.T) {
	w := ExecutionWindow{StartAt: time.Unix(100, 0), FinishAt: time.Unix(200, 0)}
	if !w.Task13WithinWorkspace(100 * time.Second) {
		t.Fatal("exact workspace limit was rejected")
	}
}
