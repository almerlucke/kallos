package tests

import (
	"testing"

	"github.com/almerlucke/kallos"
)

func TestNotes(t *testing.T) {
	major := kallos.NewScale([]int{2, 2, 1, 2, 2, 2, 1})
	triad := major.Triad(0, 60)

	t.Logf("first major triad %v\n", triad)
}
