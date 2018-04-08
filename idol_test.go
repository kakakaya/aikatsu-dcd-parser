package dcdkatsu

import (
	"testing"
)

const test_idol_id = "RXL7z1ut6wQ5sv3x"

func TestFetchIdol(t *testing.T) {
	idol, err := FetchIdol(test_idol_id)
	if err != nil {
		t.Logf("%+v", idol)
		t.Fatalf("Error occured while fetching idol.  %v.", err)
	}
}
