package compiledCheckerCC

import (
	"fmt"
	"runtime"
)

func CheckAndShow() {
	fmt.Printf("Compiled with Go version %s\n", runtime.Version())
}
