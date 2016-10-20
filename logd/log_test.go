package logd

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	SetLevel(Ldebug)

	Printf("Print: foo\n")
	Print("Print: foo")

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")

	SetLevel(Lerror)

	Printf("Print: foo\n")
	Print("Print: foo")

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")
}

func BenchmarkLogFileChan(b *testing.B) {
	log := New(LogOption{
		IsAsync:    true,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})

	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}

func BenchmarkLogFile(b *testing.B) {
	f, _ := os.OpenFile("testdata/onlyfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := New(LogOption{
		Out:        f,
		IsAsync:    false,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})

	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}
