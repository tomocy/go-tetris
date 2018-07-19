package tetris

type Field struct {
	figure Figure
	w      int
	h      int
}

func NewField(w, h int) *Field {
	f := new(Field)
	f.setUp(w, h)
	return f
}

func (f *Field) setUp(w, h int) {
	figure := make(Figure, h)
	for i := 0; i < h; i++ {
		figure[i] = make([]Level, w)
		for j := 0; j < w; j++ {
			figure[i][j] = space
		}
	}
	f.figure = figure
	f.w = w
	f.h = h
}

func (f Field) Bytes() []byte {
	b := make([]byte, 0, 10)
	for _, y := range f.figure {
		for _, x := range y {
			b = append(b, x.String()...)
		}
		b = append(b, "\n"...)
	}
	return b
}

type Level interface {
	String() string
}

type Space int
type Block int

func (s Space) String() string {
	return " "
}

func (b Block) String() string {
	return "●"
}

var (
	space Space
	block Block
)
