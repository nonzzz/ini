package tokenizer

import "fmt"

const (
	TSection    = "Section"
	TIdent      = "Ident"
	TIllgal     = "Illgal"
	TComment    = "Comment"
	TComment2   = "Comment"
	TAssign     = "Assign"
	TOpenBrace  = "OpenBrace"
	TCloseBrace = "CloseBrace"
	TKey        = "Key"
	TValue      = "Value"
	TEof        = "Eof"
)

type Tokenizer struct {
	Kind  string
	Value string
	Line  int
}

func NewToken(kind string, value string, line int) Tokenizer {

	return Tokenizer{
		Kind:  kind,
		Value: value,
		Line:  line,
	}
}

func (token Tokenizer) String() string {
	return fmt.Sprintf("{Kind:%s,Value:\"%s\" Linie:%d}",
		token.Kind,
		token.Value,
		token.Line)
}
