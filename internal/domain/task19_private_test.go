package domain

import (
	"testing"
)

func TestTask19LinkedRevisionCannotDetach(t *testing.T) {
	b := ToolRevision{ExecutionRequestID: "run-42"}
	if b.Task19CanDetach() {
		t.Fatal("linked revision was detachable")
	}
}
