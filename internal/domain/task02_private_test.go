package domain

import (
	"testing"
)

func TestTask02DisabledUserCannotAuthenticate(t *testing.T) {
	u := User{Status: UserDisabled}
	if u.Task02CanAuthenticate() {
		t.Fatal("disabled user was accepted")
	}
}
