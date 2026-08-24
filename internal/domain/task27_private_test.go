package domain

import (
	"testing"
)

func TestTask27SuspendedZoneRejectsTraffic(t *testing.T) {
	s := TrustZone{Status: TrustZoneSuspended}
	if s.Task27AcceptsTraffic() {
		t.Fatal("suspended zone accepted traffic")
	}
}
