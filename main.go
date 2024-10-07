package main

import (
	"flag"
	"fmt"
	"log"

	"os"

	"x-ui-scratch/config"
	"x-ui-scratch/logger"

	"github.com/op/go-logging"
)

func runWebServer() {
	fmt.Println("runWebServer", config.GetLogLevel())

	log.Printf("Starting %v %v", config.GetName(), config.GetVersion())

	switch config.GetLogLevel() {
	case config.Debug:
		logger.InitLogger(logging.DEBUG)
	case config.Info:
		logger.InitLogger(logging.INFO)
	case config.Notice:
		logger.InitLogger(logging.NOTICE)
	case config.Warn:
		logger.InitLogger(logging.WARNING)
	case config.Error:
		logger.InitLogger(logging.ERROR)
	default:
		log.Fatalf("Unknown log level: %v", config.GetLogLevel())
	}

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
