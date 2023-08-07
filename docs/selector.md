# Selector


## Interface

```go

const (
	SectionKind    = ast.KSection
	ExpressionKind = ast.KExpression
	CommentKind    = ast.KComment
)

type Operate interface {
	Get() (ast.Element, error)
	Set(bindings AttributeBindings) bool
	Delete() bool
	InsertBefore(node ast.Element) bool
	InsertAfter(node ast.Element) bool
}

type Selector interface {
	Section(id string) Operate
	Comment(id string) Operate
	Expression(key string) Operate
}

```

## Usage


```go

i := ini.New()
i.loadFile("./test.ini")

selector := ini.NewSelector(i)

```