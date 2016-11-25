package log

import "testing"

func TestLevelComparison(t *testing.T) {
	if !(NOTSET < DEBUG &&
		DEBUG < INFO &&
		INFO < WARN &&
		WARN < FATA) {
		t.Fatal("Wrong relationship between levels")
	}
}
