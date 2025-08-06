package fields

// Field is one of the rendered fields. If inline is set to true, both
// the label and the value will be rendered on the same line.
type Field struct {
	Label  string
	Value  string
	Inline bool
}
