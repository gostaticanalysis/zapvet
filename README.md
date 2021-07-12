# zapvet

[![pkg.go.dev][gopkg-badge]][gopkg]

`zapvet` is static analysis tool for [zap](https://pkg.go.dev/go.uber.org/zap).

* [fieldtype](./passes/fieldtype): fieldtype finds confliction type of field

## Install

You can get `zapvet` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/gostaticanalysis/zapvet@latest
```

## How to use

`zapvet` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which zapvet) ./...
```

## Analyzers

### fieldtype

[fieldtype](./passes/fieldtype) finds confliction type of field.

```go
package a

import "go.uber.org/zap"

func f() {
	zap.String("id", "100")
	zap.Int("id", 100)       // want `"id" conflict type Int vs String`
	zap.Any("id", "100")     // OK - ignore
	zap.Reflect("id", "100") // OK - ignore
	zap.String("xxx", "100") // OK
}
```

## Analyze with golang.org/x/tools/go/analysis

You can get analyzers of zapvet from [zapvet.Analyzers](https://pkg.go.dev/github.com/gostaticanalysis/zapvet/#Analyzers).
And you can use them with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/zapvet
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/zapvet?status.svg
