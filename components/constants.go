package rephtml

type Strictness string

const (
	DEFAULT Strictness = "default"
	LAZY    Strictness = "lazy"
	STRICT  Strictness = "strict"
)
