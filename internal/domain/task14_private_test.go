package domain

import (
	"testing"
)

func TestTask14EqualRiskBoundsAreInvalid(t *testing.T) {
	r := RiskRange{Minimum: 5000, Maximum: 5000}
	if r.Task14ValidForPolicy() {
		t.Fatal("equal risk bounds were accepted")
	}
}
