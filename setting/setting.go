package setting

func init() {
	Field.setUp()
}

var Field field

type field struct {
	w int
	h int
}

func (f *field) setUp() {
	f.w = 10
	f.h = 20
}

func (f field) Width() int {
	return f.w
}

func (f field) Height() int {
	return f.h
}
