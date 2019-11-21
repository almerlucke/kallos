package gokallos

// TransformType identifies the type of transform
type TransformType int

const (
	// TransformDuration apply transform only on duration
	TransformDuration TransformType = iota
	// TransformPitch apply transform only on pitch
	TransformPitch
	// TransformVelocity apply transform only on velocity
	TransformVelocity
	// TransformChannel apply transform only on channel
	TransformChannel
	// TransformAllButChannel apply transform on all fields except channel
	TransformAllButChannel
	// TransformAll apply transform on all fields
	TransformAll
)

// Transformer can transform a stream, should always return a new stream
type Transformer interface {
	TransformValue(v Value) Value
	Type() TransformType
}

// basicTransformer implements the basic transformer methods
type basicTransformer struct {
	transformType TransformType
}

// TransformType returns the type of transform
func (t *basicTransformer) TransformType() TransformType {
	return t.transformType
}
