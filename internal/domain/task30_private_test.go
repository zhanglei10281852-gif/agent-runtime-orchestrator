package domain

import (
	"testing"
)

func TestTask30EqualRiskBoundsAreInvalid(t *testing.T) {
	r := RiskRange{Minimum: 5000, Maximum: 5000}
	if r.Task30ValidForPolicy() {
		t.Fatal("equal risk bounds were accepted")
	}
}
