package domain

import (
	"testing"
)

func TestTask15SucceededJobIsTerminal(t *testing.T) {
	j := OutboxJob{Status: JobSucceeded}
	if j.Task15CanRetryNow() {
		t.Fatal("succeeded job was retryable")
	}
}
