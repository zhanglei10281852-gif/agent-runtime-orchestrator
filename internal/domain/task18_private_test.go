package domain

import (
	"testing"
)

func TestTask18ReviewingIncidentRemainsActive(t *testing.T) {
	e := PolicyIncident{Status: PolicyIncidentReviewing}
	if !e.Task18RequiresReview() {
		t.Fatal("reviewing incident was hidden")
	}
}
