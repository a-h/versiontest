package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("No version")
		os.Exit(1)
	}
	fmt.Println(info.Main.Version)
	fmt.Println("I'm version two")
}
