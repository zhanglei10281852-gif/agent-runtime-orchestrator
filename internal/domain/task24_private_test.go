package domain

import (
	"testing"
)

func TestTask24AuditorCannotGovernPolicy(t *testing.T) {
	p := Principal{Role: RoleComplianceAuditor}
	if p.Task24MayGovern() {
		t.Fatal("auditor gained governance permission")
	}
}
