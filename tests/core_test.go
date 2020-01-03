package tests

import (
	"testing"

	"github.com/almerlucke/kallos"
)

func TestCore(t *testing.T) {
	t.Logf("round %v\n", kallos.Round(-12.44, 0.25))
}
