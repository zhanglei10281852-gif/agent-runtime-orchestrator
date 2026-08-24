package domain

import (
	"testing"
)

func TestTask11SuspendedZoneRejectsTraffic(t *testing.T) {
	s := TrustZone{Status: TrustZoneSuspended}
	if s.Task11AcceptsTraffic() {
		t.Fatal("suspended zone accepted traffic")
	}
}
