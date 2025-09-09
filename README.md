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
goos: linux
goarch: amd64
pkg: github.com/tevino/log
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkLog15-4         	  528672	      2178 ns/op
BenchmarkLogrus-4        	  502843	      2170 ns/op
BenchmarkLoggo-4         	  537241	      2214 ns/op
BenchmarkGoLogging-4     	 1807219	       640.1 ns/op
BenchmarkThisPackage-4   	 3868623	       309.8 ns/op 👈
BenchmarkStdLog-4        	 4278027	       274.7 ns/op
```

go 1.24.0

| Library | Version |
|-------|--------|
| log15 | v3.0.0 |
| logrus | v1.9.3 |
| loggo | v1.0.0 |
| go-logging | v0.0.0-20160315200505-970db520ece7 |

## Acknowledgements

- @kirk91: Provide APIs to make it possible for user to specify a custom caller offset.
