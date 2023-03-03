package lexer

type Lexer struct {
	Input    []byte
	char     byte
	pos      int
	read_pos int
	line     int
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{Input: []byte(input)}

	return lexer
}
