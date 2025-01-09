# README

Create tools.go with the content:

```
//go:build tools
// +build toos

package tools

import (
	_"github.com/99designs/gqlgen"
	_"github.com/99designs/gqlgen/graphql/introspection"
)
```

Download dependencies of `tools.go` file:
```
$ go run github.com/99designs/gqlgen init
```

Start server:
```
$ go run server.go
```

## Troubleshooting
#### In case of error:
go: modules disabled by GO111MODULE=off; see 'go help modules'

#### Solving:
Export `GO111MODULE`with value `on`.
```
$ export GO111MODULE="on"
```
#### Missing dependencies
Execute the command:
```
$ go mod tidy
```