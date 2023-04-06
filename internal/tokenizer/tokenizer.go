package tokenizer

type T uint8

const (
	TEof T = iota
	TSection
	TComment
	TExpression
	TAssign
	TKey
	TValue
	TDocument
)

var tokenToString = []string{
	"EndOfFile",
	"SectionDeclaration",
	"CommentDeclaration",
	"ExpressionDeclaration",
}

func (t T) String() string {
	return tokenToString[t]
}
