package domain

import (
	"testing"
)

func TestTask09ReviewingIncidentRemainsActive(t *testing.T) {
	e := PolicyIncident{Status: PolicyIncidentReviewing}
	if !e.Task09RequiresReview() {
		t.Fatal("reviewing incident was hidden")
	}
}
