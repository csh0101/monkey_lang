package lexer

type Lexer struct {
	input        string
	position     int // the d
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
	l.position++
	l.readPosition++
}
