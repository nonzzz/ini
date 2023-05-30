# Ini

<p align="center">
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/nonzzz/ini"><img src="https://goreportcard.com/badge/github.com/nonzzz/ini?style=flat-square" /></a>
<a title="Doc for grm" target="_blank" href="https://pkg.go.dev/github.com/nonzzz/ini"><img src="https://pkg.go.dev/badge/github.com/nonzzz/ini.svg" /></a>
<a title="Codecov" target="_blank" href="https://codecov.io/gh/nonzzz/ini"><img src="https://img.shields.io/codecov/c/github/nonzzz/ini?style=flat-square&logo=codecov" /></a>
</p>

A simple standard ini parser with golang.

## Install

```bash

$ go get github.com/nonzzz/ini

```

## Features

- Read by file.
- Read by string.
- Marshal to Json or Map.
- Support Accpect (visitor pattern).

## Usage

```go

i := ini.New()

//  Load File
i.LoadFile("your ini file")

// Parse

txt :=`

[s]

a = 3

[s1]

b = 4

`

ini.Parse(txt)

```

## Abstract syntax tree

```

Node {
    Type: doc,
    Nodes: [
        Node {
            Type: section
            Nodes: []
        },
        Node {
            Type: section
            Nodes: []
        },
    ]
    Loc: {
        Start: 0,
        Len: 0,
    }

}

```

```go


import "github.com/nonzzz/ini/pkg/ast"

type visitor struct {
    ini.IniVisitor
}

func (v *visitor) Section(node *ast.Section) {}

func (v *visitor) Expression(node *ast.Expression) {}

func (v *visitor) Comment(node *ast.Comment) {}

```

## Acknowledgements

Thanks to [JetBrains](https://www.jetbrains.com/) for allocating free open-source licences for IDEs.

<p align="left">
<img width="250px" height="250px"  src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand_icon.png" alt="GoLand logo.">
</p>

## LICENSE

[MIT](LICENSE)
