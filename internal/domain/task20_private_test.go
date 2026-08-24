package domain

import (
	"testing"
)

func TestTask20ExecutingRequestCannotCancel(t *testing.T) {
	r := ExecutionRequest{State: ExecutionRequestExecuting}
	if r.Task20CanCancel() {
		t.Fatal("executing request was cancellable")
	}
}
