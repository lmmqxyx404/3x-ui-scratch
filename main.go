package main

import (
	"flag"
	"fmt"
	"log"

	"os"

	"x-ui-scratch/config"
	"x-ui-scratch/database"
	"x-ui-scratch/logger"
	"x-ui-scratch/sub"
	"x-ui-scratch/web"
	"x-ui-scratch/web/global"

	"github.com/op/go-logging"
)

func runWebServer() {
	fmt.Println("runWebServer", config.GetLogLevel())

	log.Printf("Starting %v %v", config.GetName(), config.GetVersion())
	// set the logger
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

	err := database.InitDB(config.GetDBPath())
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	var server *web.Server
	server = web.NewServer()
	// log.Fatalf("Error starting web server: %v", server)
	global.SetWebServer(server)

	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting web server: %v", err)
		return
	}

	var subServer *sub.Server
	subServer = sub.NewServer()
	global.SetSubServer(subServer)
	err = subServer.Start()
	if err != nil {
		log.Fatalf("Error starting sub server: %v", err)
		return
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
