package main

import(
	"fmt"
	"runtime"
)

func getOs() string {
	return runtime.GOOS
}

func main() {
	fmt.Println(getOs())
}
