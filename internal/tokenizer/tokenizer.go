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
