package tests

import (
	"github.com/almerlucke/kallos/notes"
	"testing"
)

func TestNotes(t *testing.T) {
	major := notes.NewScale([]int{2, 2, 1, 2, 2, 2, 1})

	triad1 := major.Triad(1)
	triad2 := triad1.Invert(1)
	triad3 := triad1.Invert(2)
	triad4 := triad1.Invert(3)

	c := notes.Chord{2, 4, 6}.SnapToScale(major, 1).Deduplicate()

	t.Logf("first major triad %v\n", triad1)
	t.Logf("invert1 triad %v\n", triad2)
	t.Logf("invert2 triad %v\n", triad3)
	t.Logf("invert3 triad %v\n", triad4)
	t.Logf("snap1 %v\n", c)
}
