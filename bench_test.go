package log

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/juju/loggo"
	gologging "github.com/op/go-logging"
	"github.com/sirupsen/logrus"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}
}

func BenchmarkLog15(b *testing.B) {
	var buf bytes.Buffer
	var logger = log15.New("test", "test")
	logger.SetHandler(log15.StreamHandler(&buf, log15.LogfmtFormat()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test: ", hostname, nil)
	}
}

func BenchmarkLogrus(b *testing.B) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logrus.WithFields(logrus.Fields{
			"hostname": hostname,
		}).Info("This is a test: ")
	}
}

func BenchmarkLoggo(b *testing.B) {
	var buf bytes.Buffer
	var logger = loggo.GetLogger("")
	logger.SetLogLevel(loggo.INFO)
	if _, err := loggo.RemoveWriter("default"); err != nil {
		panic(err)
	}
	var writer = loggo.NewColorWriter(&buf)
	if err := loggo.RegisterWriter("default", writer); err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a test: %s", hostname)
	}
}

func BenchmarkGoLogging(b *testing.B) {
	var buf bytes.Buffer
	var logger = gologging.MustGetLogger("test")
	var backend = gologging.NewLogBackend(&buf, "test", 0)
	leveledBackend := gologging.AddModuleLevel(backend)
	logger.SetBackend(leveledBackend)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a test: %s", hostname)
	}
}

func BenchmarkThisPackage(b *testing.B) {
	var buf bytes.Buffer
	var logger = NewLeveledLogger(&buf, LstdFlags)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Printf("This is a test: %s", hostname)
	}
}

func BenchmarkStdLog(b *testing.B) {
	var buf bytes.Buffer
	var logger = log.New(&buf, "test", log.LstdFlags)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Printf("This is a test: %s", hostname)
	}
}
