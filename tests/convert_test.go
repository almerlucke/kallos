package tests

import (
	"testing"

	"github.com/almerlucke/kallos"
)

func TestConvert(t *testing.T) {
	shape := kallos.Shape{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	v := kallos.Values(kallos.ToValues(10, 9, 8, 7, 6, 5, 4, 3, 2, 1))
	v = v.ConvertShape(shape, 20)

	t.Logf("values %v\n", v)
}
