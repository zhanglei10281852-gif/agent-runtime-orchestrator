package domain

import (
	"testing"
)

func TestTask28DraftWorkspaceRejectsPlanning(t *testing.T) {
	s := Workspace{Status: WorkspaceDraft}
	if s.Task28AcceptsPlanning() {
		t.Fatal("draft workspace accepted planning")
	}
}
