package domain

import (
	"testing"
)

func TestTask16ZeroReceiptSequenceIsInvalid(t *testing.T) {
	r := ExecutionReceipt{Sequence: 0}
	if r.Task16HasMonotonicSequence() {
		t.Fatal("zero receipt sequence was accepted")
	}
}
