package generators

import "github.com/almerlucke/kallos"

// Chain generators together one after the other
type Chain struct {
	chain      []kallos.Generator
	index      int
	current    kallos.Generator
	continuous bool
	done       bool
}

func NewChain(chain ...kallos.Generator) *Chain {
	return &Chain{
		chain:      chain,
		index:      0,
		current:    chain[0],
		continuous: true,
		done:       false,
	}
}

func (c *Chain) SetContinuous(continuous bool) {
	c.continuous = continuous
}

// GenerateValue by getting next value from current generator, if current is continuous we
// only get one value, otherwise we deplete the current generator before going to the next
// one.
func (c *Chain) GenerateValue() (value kallos.Value) {
	gotoNext := false

	value = c.current.GenerateValue()

	if c.current.IsContinuous() {
		gotoNext = true
	} else if c.current.Done() {
		gotoNext = true
		if c.IsContinuous() {
			c.current.Reset()
		}
	}

	if gotoNext {
		c.index += 1
		if c.index >= len(c.chain) {
			c.done = !c.continuous
			c.index = 0
		}
		c.current = c.chain[c.index]
	}

	return
}

func (c *Chain) IsContinuous() bool {
	return c.continuous
}

func (c *Chain) Done() bool {
	return c.done
}

// Reset all generators
func (c *Chain) Reset() {
	c.index = 0
	c.done = false
	for _, gen := range c.chain {
		gen.Reset()
	}
}
