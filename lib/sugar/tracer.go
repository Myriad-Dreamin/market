package sugar

import (
	"fmt"
	"runtime"
)

func PrintStack() {
	var buf [1024 * 10]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}
