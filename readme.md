a plugin for `golangci-lint` to disallow using the `fmt` package to print.

build with:
```
go build -buildmode=plugin plugin/nofmtprintf.go
```

test with:
```
go test ./...
```