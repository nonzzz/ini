package tokenizer

type T uint8

const (
	TEof T = iota
	TSection
	TComment
	TAssign
	TKey
	TValue
	TDocument
)
