package generators

import (
	"math/rand"

	"github.com/almerlucke/kallos"
)

// RandomWalkMatrix any dimension matrix for random walk
type RandomWalkMatrix interface {
	ValueForIndex(index []int) kallos.Value
}

// RandomWalk1DMatrix for 1D matrix
type RandomWalk1DMatrix struct {
	Values kallos.Values
}

// ValueForIndex for 1D matrix
func (r *RandomWalk1DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]]
}

// RandomWalk2DMatrix for 2D matrix
type RandomWalk2DMatrix struct {
	Values []kallos.Values
}

// ValueForIndex for 2D matrix
func (r *RandomWalk2DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]][index[1]]
}

// RandomWalk3DMatrix for 3D matrix
type RandomWalk3DMatrix struct {
	Values [][]kallos.Values
}

// ValueForIndex for 3D matrix
func (r *RandomWalk3DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]][index[1]][index[2]]
}

// RandomWalk structure
type RandomWalk struct {
	Dimensions []int
	Index      []int
	Matrix     RandomWalkMatrix
}

// NewRandomWalk creates a new random walk
func NewRandomWalk(dimensions []int, matrix RandomWalkMatrix) *RandomWalk {
	rw := &RandomWalk{
		Dimensions: dimensions,
		Matrix:     matrix,
	}

	index := make([]int, len(dimensions))

	for i, d := range dimensions {
		index[i] = rand.Intn(d)
	}

	rw.Index = index

	return rw
}

// GenerateValue for random walk
func (r *RandomWalk) GenerateValue() kallos.Value {
	selectedDimension := rand.Intn(len(r.Dimensions))
	selectedDimensionSize := r.Dimensions[selectedDimension]
	selectedIndex := r.Index[selectedDimension]

	if selectedIndex < 1 {
		selectedIndex++
	} else if selectedIndex > (selectedDimensionSize - 2) {
		selectedIndex--
	} else {
		if rand.Intn(2) == 1 {
			selectedIndex++
		} else {
			selectedIndex--
		}
	}

	r.Index[selectedDimension] = selectedIndex

	return r.Matrix.ValueForIndex(r.Index)
}

// IsContinuous for random walk
func (r *RandomWalk) IsContinuous() bool {
	return true
}

// Done for random walk
func (r *RandomWalk) Done() bool {
	return false
}

// Reset for random walk
func (r *RandomWalk) Reset() {
}
