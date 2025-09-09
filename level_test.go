package log

import "testing"

func TestLevelComparison(t *testing.T) {
	levelsSortedByValueAsc := []Level{NOTSET, DEBUG, INFO, WARN, FATA}
	for i, level := range levelsSortedByValueAsc[:len(levelsSortedByValueAsc)-1] {
		if level >= levelsSortedByValueAsc[i+1] {
			t.Errorf("Expected level %v to be less than %v", level, levelsSortedByValueAsc[i+1])
		}
	}
}
