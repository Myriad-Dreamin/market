package abstract_test

type ContextHelperInterface interface {
	Helper()
	Log(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Logf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
