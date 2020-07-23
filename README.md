# go-load

benchmarks on loading files methods

## Test loading methods

The `main.go` file aims just to test both methods

```sh
go run .\main.go -testfile C:/Users/damien/Documents/workspace/go-load/mock/sample2K.log
```

## Benchmarks

```sh
go test -bench="." -benchmem ./... -args -path=C:/Users/damien/Documents/workspace/go-load/mock
```