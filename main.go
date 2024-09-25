package main

import (
	"flag"
	"fmt"

	"os"
)

func runWebServer() {
	fmt.Println("runWebServer")

	// log.Printf("Starting %v %v", config.GetName(), config.GetVersion())
}

func main() {
	if len(os.Args) < 2 {
		runWebServer()
		return
	}

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	fmt.Println("hello world")
}
