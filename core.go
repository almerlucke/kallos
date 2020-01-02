package kallos

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	// TicksPerQuarterNote constant
	TicksPerQuarterNote float64 = 192.0
	// BPM constant
	BPM float64 = 120.0
	// SecondsPerBeat constant
	SecondsPerBeat float64 = 0.5
	// BeatsPerSecond constant
	BeatsPerSecond float64 = 2.0
)

var baseNoteNameMap = map[string]int{
	"c":  0,
	"c#": 1,
	"db": 1,
	"d":  2,
	"d#": 3,
	"eb": 3,
	"e":  4,
	"f":  5,
	"f#": 6,
	"gb": 6,
	"g":  7,
	"g#": 8,
	"ab": 8,
	"a":  9,
	"a#": 10,
	"bb": 10,
	"b":  11,
}

var baseNotePrefixList = []string{
	"c#",
	"c",
	"db",
	"d#",
	"d",
	"eb",
	"e",
	"f#",
	"f",
	"gb",
	"g#",
	"g",
	"ab",
	"a#",
	"a",
	"bb",
	"b",
}

// Value wraps a generated value slice
type Value []float64

// Values wraps a slice of Value's
type Values []Value

// NoteNumberFromNoteName convert a note name to midi note number
func NoteNumberFromNoteName(name string) (int, error) {
	name = strings.ToLower(name)
	foundPrefix := ""

	for _, prefix := range baseNotePrefixList {
		if strings.HasPrefix(name, prefix) {
			foundPrefix = prefix
			break
		}
	}

	if foundPrefix == "" {
		return 0, errors.New("invalid note number")
	}

	offset, err := strconv.Atoi(name[len(foundPrefix):])
	if err != nil {
		return 0, err
	}

	return 12 + offset*12 + baseNoteNameMap[foundPrefix], nil
}

// Clip value to low and high
func Clip(v float64, low float64, high float64) float64 {
	return math.Min(math.Max(float64(v), float64(low)), float64(high))
}

// ToValues convert a number of floats to a slice of Values
func ToValues(v ...float64) Values {
	vs := make(Values, len(v))

	for i, vc := range v {
		vs[i] = Value{vc}
	}

	return vs
}
