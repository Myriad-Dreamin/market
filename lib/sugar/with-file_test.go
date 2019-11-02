package sugar

import (
	"errors"
	"os"
	"testing"
)

func TestWithFile(t *testing.T) {
	WithFile(func(f *os.File) {
		if f.Name() != "test.log" {
			t.Error(errors.New("open failed"))
		}
	}, "test.log")
	_ = os.Remove("test.log")
}

func TestWithFiles(t *testing.T) {
	var log1, log2 = new(os.File), new(os.File)
	WithFiles(func() {
		if log1 == nil || log2 == nil {
			t.Error(errors.New("open failed"), log1 == nil, log2 == nil)
		}
		if log1.Name() != "test1.log" {
			t.Error(errors.New("open failed"), log1)
		}
		if log2.Name() != "test2.log" {
			t.Error(errors.New("open failed"), log2)
		}
		_, err := log1.WriteString("test1.log")
		if err != nil {
			t.Error(errors.New("bad opened test1.log"), err)
		}
		_, err = log2.WriteString("test2.log")
		if err != nil {
			t.Error(errors.New("bad opened test2.log"), err)
		}
	}, Files(log1, log2), "test1.log", "test2.log")
	_ = os.Remove("test1.log")
	_ = os.Remove("test2.log")
}
