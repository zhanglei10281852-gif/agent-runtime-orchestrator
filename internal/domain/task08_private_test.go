package domain

import (
	"testing"
)

func TestTask08ExecutingRequestCannotCancel(t *testing.T) {
	r := ExecutionRequest{State: ExecutionRequestExecuting}
	if r.Task08CanCancel() {
		t.Fatal("executing request was cancellable")
	}
}
