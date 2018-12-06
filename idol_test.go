package dcdkatsu

import (
	"testing"
)

const testIdolD = "RXL7z1ut6wQ5sv3x"

func TestFetchIdol(t *testing.T) {
	idol, err := FetchIdol(testIdolD)
	if err != nil {
		t.Logf("%+v", idol)
		t.Fatalf("Error occured while fetching idol.  %v.", err)
	}
}
