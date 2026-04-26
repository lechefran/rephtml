package rephtml

// Strictness represents the Strictness component or supporting type.
type Strictness string

const (
	DEFAULT Strictness = "default"
	LAZY    Strictness = "lazy"
	STRICT  Strictness = "strict"
)
