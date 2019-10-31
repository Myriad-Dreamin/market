package abstract_test

import "log"

type ContextHelper struct {
	*log.Logger
}

func (h ContextHelper) Helper() {
}

func (h ContextHelper) Log(args ...interface{}) {
	h.Logger.Println(args...)
}

func (h ContextHelper) Error(args ...interface{}) {
	h.Logger.Println(args...)
}

func (h ContextHelper) Logf(format string, args ...interface{}) {
	h.Logger.Printf(format, args...)
}

func (h ContextHelper) Errorf(format string, args ...interface{}) {
	h.Logger.Printf(format, args...)
}
