package domain

import (
	"testing"
)

func TestTask07LinkedRevisionCannotDetach(t *testing.T) {
	b := ToolRevision{ExecutionRequestID: "run-42"}
	if b.Task07CanDetach() {
		t.Fatal("linked revision was detachable")
	}
}
