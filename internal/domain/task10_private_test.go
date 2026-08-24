package domain

import (
	"testing"
)

func TestTask10OpenIncidentBlocksArchive(t *testing.T) {
	r := RunReadiness{ExecutionRequestState: ExecutionRequestCompleted, OpenPolicyIncident: true}
	if r.Task10CompleteForArchive() {
		t.Fatal("open incident did not block archive")
	}
}
