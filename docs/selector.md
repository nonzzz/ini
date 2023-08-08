# Selector


## Interface

```go

const (
	SectionKind    = ast.KSection
	ExpressionKind = ast.KExpression
	CommentKind    = ast.KComment
)

type AttributeBindings struct {
	Id    string
	Text  string
	Key   string
	Value string
}

type Operate interface {
	Get() (ast.Element, error)
	Set(bindings AttributeBindings) bool
	Delete() bool

type Selector interface {
	Query(id string, kind ast.K) Operate
}

```

## Usage


```go

i := ini.New()
i.loadFile("./test.ini")

selector := ini.NewSelector(i)

```