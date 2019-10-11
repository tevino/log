# log
A lightweight, leveled logging package for Go.


## Install
`go get github.com/tevino/log`


## Features
- Leveled logging and filtering
- Colored output
- Log with file name and line number

## Example

```go
import "github.com/tevino/log"

log.SetOutputLevel(log.INFO)

log.Debugf("Output level is %s.", log.OutputLevel())
log.Info("Thus only this line is printed.")
```


## Benchmark
```
BenchmarkLog15-4              300000          4391 ns/op        1202 B/op         16 allocs/op
BenchmarkLogrus-4             500000          3444 ns/op        1715 B/op         24 allocs/op
BenchmarkLoggo-4              500000          3275 ns/op         344 B/op         11 allocs/op
BenchmarkGoLogging-4         1000000          1136 ns/op         601 B/op         10 allocs/op
BenchmarkThisPackage-4       3000000           731 ns/op         214 B/op          2 allocs/op
BenchmarkStdLog-4            2000000           571 ns/op         174 B/op          2 allocs/op
```

## Acknowledgements

- @kirk91: Provide APIs to make it possible for user to specify a custom caller offset.
