package main

import(
	"fmt"
	"runtime"
)

// get the current os that go is installed on
func getOs() string {
	return runtime.GOOS
}

func main() {
	InitCliArgs()
	fmt.Println(getOs())
}
