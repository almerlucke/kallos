package kallos

import (
	"errors"
	"strconv"
	"strings"
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

type Scale struct {
	steps   []int
	indices []int
}

func NewScale(steps []int) *Scale {
	indices := make([]int, len(steps))
	index := 0

	for i, s := range steps {
		indices[i] = index
		index += s
	}

	return &Scale{
		steps:   steps,
		indices: indices,
	}
}

func (s Scale) Triad(index int, root int) []int {
	triad := make([]int, 3)

	triad[0] = root + s.indices[index]

	pos := index + 2
	if pos >= len(s.indices) {
		triad[1] = root + s.indices[pos%len(s.indices)] + 12
	} else {
		triad[1] = root + s.indices[pos]
	}

	pos = index + 4
	if pos >= len(s.indices) {
		triad[2] = root + s.indices[pos%len(s.indices)] + 12
	} else {
		triad[2] = root + s.indices[pos]
	}

	return triad
}

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
