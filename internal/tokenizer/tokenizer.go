package tokenizer

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
