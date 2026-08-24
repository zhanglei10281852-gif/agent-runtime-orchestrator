package domain

import (
	"testing"
)

func TestTask23DisabledUserCannotAuthenticate(t *testing.T) {
	u := User{Status: UserDisabled}
	if u.Task23CanAuthenticate() {
		t.Fatal("disabled user was accepted")
	}
}
