package domain

import (
	"testing"
)

func TestTask12DraftWorkspaceRejectsPlanning(t *testing.T) {
	s := Workspace{Status: WorkspaceDraft}
	if s.Task12AcceptsPlanning() {
		t.Fatal("draft workspace accepted planning")
	}
}
