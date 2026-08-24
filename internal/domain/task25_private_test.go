package domain

import (
	"testing"
)

func TestTask25ReviewingIncidentRemainsActive(t *testing.T) {
	e := PolicyIncident{Status: PolicyIncidentReviewing}
	if !e.Task25RequiresReview() {
		t.Fatal("reviewing incident was hidden")
	}
}
