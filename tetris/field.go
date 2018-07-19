package tetris

type Field [][]Level

func MakeField(w, h int) Field {
	f := make(Field, h)
	for i := 0; i < len(f); i++ {
		f[i] = make([]Level, w)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = space
		}
	}

	return f
}

func (f Field) Bytes() []byte {
	b := make([]byte, 0, 10)
	for _, y := range f {
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
