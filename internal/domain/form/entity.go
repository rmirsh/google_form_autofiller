package form

type Field struct {
	ID        string
	Container string
	TypeID    float64
	Required  bool
	Options   []string
}

type Form struct {
	Fields []Field
}
