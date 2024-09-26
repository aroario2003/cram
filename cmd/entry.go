package cram

import(
	"fmt"
	"runtime"
)

// get the current os that go is installed on
func GetOs() string {
	return runtime.GOOS
}

func Entry() {
	InitCliArgs()
	fmt.Println(GetOs())
}
