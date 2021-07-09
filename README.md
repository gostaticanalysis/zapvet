# zapvet

[![pkg.go.dev][gopkg-badge]][gopkg]

`zapvet` is static analysis tool for [zap](https://pkg.go.dev/go.uber.org/zap).

* [fieldtype](./passes/fieldtype): fieldtype finds confliction type of field

```sh
$ go install https://github.com/gostaticanalysis/zapvet@latest
$ go vet -vettool=`zapvet` ./...
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/zapvet
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/zapvet?status.svg
