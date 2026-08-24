package domain

import (
	"testing"
)

func TestTask26OpenIncidentBlocksArchive(t *testing.T) {
	r := RunReadiness{ExecutionRequestState: ExecutionRequestCompleted, OpenPolicyIncident: true}
	if r.Task26CompleteForArchive() {
		t.Fatal("open incident did not block archive")
	}
}
