package domain

import (
	"testing"
)

func TestTask03AuditorCannotGovernPolicy(t *testing.T) {
	p := Principal{Role: RoleComplianceAuditor}
	if p.Task03MayGovern() {
		t.Fatal("auditor gained governance permission")
	}
}
