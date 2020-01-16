package kallos

import (
	"sort"
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

// Chord is a slice of indices of a scale
type Chord []int

// Scale represents the steps in a note scale
type Scale struct {
	steps   []int
	indices []int
}

// NewChord creates a new chord
func NewChord(size int) Chord {
	return make(Chord, size)
}

// Len for sorting
func (c Chord) Len() int {
	return len(c)
}

// Swap for sorting
func (c Chord) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less for sorting
func (c Chord) Less(i, j int) bool {
	return c[i] < c[j]
}

// Copy a chord
func (c Chord) Copy() Chord {
	cc := NewChord(len(c))

	for i, n := range c {
		cc[i] = n
	}

	return cc
}

// SnapToScale snap the chord to fit the scale, if direction > 0 we round to the higher note if
// there are two scale notes that have the same distance to the given note, returns a new chord.
func (c Chord) SnapToScale(sc *Scale, direction int) Chord {
	cc := c.Copy()

	for i, n := range c {
		cc[i] = sc.SnapNote(n, direction)
	}

	return cc
}

// Transpose the chord by a certain amount, returns a new chord
func (c Chord) Transpose(amount int) Chord {
	cc := c.Copy()

	for i, n := range c {
		cc[i] = n + amount
	}

	return cc
}

// Deduplicate removes duplicate notes
func (c Chord) Deduplicate() Chord {
	uniqueNotes := map[int]int{}

	for _, n := range c {
		uniqueNotes[n] = 1
	}

	cc := Chord{}
	for k := range uniqueNotes {
		cc = append(cc, k)
	}

	return cc.Sort()
}

// Normalize a chord by lowering octaves, returns a new chord
func (c Chord) Normalize() Chord {
	cc := c.Copy()

	min := 127

	// first get minimum note
	for _, n := range cc {
		if n < min {
			min = n
		}
	}

	// lower octaves if minimum note is higher then 1 octave
	if min >= 12 {
		// round to nearest octave
		octaves := min / 12 * 12
		// subtract octaves from notes
		for i, n := range cc {
			cc[i] = n - octaves
		}
	}

	return cc
}

// Sort a chord, returns a new chord
func (c Chord) Sort() Chord {
	cc := c.Copy()

	sort.Sort(cc)

	return cc
}

// Invert a chord, returns a new chord
func (c Chord) Invert(times int) Chord {
	cc := c.Copy()

	for times > 0 {
		times--
		cc[0] += 12
		cc = cc.Sort().Normalize()
	}

	return cc
}

// ToValue converts a chord to a value and adds a root note
func (c Chord) ToValue(root int) Value {
	v := make(Value, len(c))

	for i, n := range c {
		v[i] = float64(root + n)
	}

	return v
}

// NewScale steps
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

// MajorScale returns a major scale
func MajorScale() *Scale {
	return NewScale([]int{2, 2, 1, 2, 2, 2, 1})
}

// MinorScale returns a minor scale
func MinorScale() *Scale {
	return NewScale([]int{2, 1, 2, 2, 2, 1, 2})
}

// MajorPentatonicScale return a major pentatonic scale
func MajorPentatonicScale() *Scale {
	return NewScale([]int{2, 2, 3, 2})
}

// MinorPentatonicScale return a minor pentatonic scale
func MinorPentatonicScale() *Scale {
	return NewScale([]int{3, 2, 2, 3})
}

// BluesScale returns a blues scale
func BluesScale() *Scale {
	return NewScale([]int{3, 2, 1, 1, 3, 2})
}

// SnapNote snaps a note to the scale
func (s Scale) SnapNote(note int, direction int) int {
	octaves := note / 12
	normalizedNote := note - octaves*12
	nearestIndex := 0
	minDistance := 12

	for _, i := range s.indices {
		distance := i - normalizedNote
		if distance < 0 {
			distance *= -1
		}

		if direction > 0 {
			if distance <= minDistance {
				minDistance = distance
				nearestIndex = i
			}
		} else {
			if distance < minDistance {
				minDistance = distance
				nearestIndex = i
			}
		}
	}

	return nearestIndex + octaves*12
}

// Triad returns a three note chord from an index in the circle of thirds
func (s Scale) Triad(index int) Chord {
	triad := NewChord(3)

	triad[0] = s.indices[index]

	pos := index + 2
	if pos >= len(s.indices) {
		triad[1] = s.indices[pos%len(s.indices)] + 12
	} else {
		triad[1] = s.indices[pos]
	}

	pos = index + 4
	if pos >= len(s.indices) {
		triad[2] = s.indices[pos%len(s.indices)] + 12
	} else {
		triad[2] = s.indices[pos]
	}

	return triad
}

// Seventh returns a four note seventh chord from an index in the circle of thirds
func (s Scale) Seventh(index int) Chord {
	seventh := NewChord(4)

	seventh[0] = s.indices[index]

	pos := index + 2
	if pos >= len(s.indices) {
		seventh[1] = s.indices[pos%len(s.indices)] + 12
	} else {
		seventh[1] = s.indices[pos]
	}

	pos = index + 4
	if pos >= len(s.indices) {
		seventh[2] = s.indices[pos%len(s.indices)] + 12
	} else {
		seventh[2] = s.indices[pos]
	}

	pos = index + 6
	if pos >= len(s.indices) {
		seventh[3] = s.indices[pos%len(s.indices)] + 12
	} else {
		seventh[3] = s.indices[pos]
	}

	return seventh
}

// NoteByName convert a note name to midi note number
func NoteByName(name string) int {
	name = strings.ToLower(name)
	foundPrefix := ""

	for _, prefix := range baseNotePrefixList {
		if strings.HasPrefix(name, prefix) {
			foundPrefix = prefix
			break
		}
	}

	if foundPrefix == "" {
		return 0
	}

	offset, err := strconv.Atoi(name[len(foundPrefix):])
	if err != nil {
		return 0
	}

	return 12 + offset*12 + baseNoteNameMap[foundPrefix]
}
