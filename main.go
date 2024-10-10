package main

import (
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

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

	// server *web.Server
	var server = web.NewServer()
	// log.Fatalf("Error starting web server: %v", server)
	global.SetWebServer(server)

	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting web server: %v", err)
		return
	}

	// subServer *sub.Server
	var subServer = sub.NewServer()
	global.SetSubServer(subServer)
	err = subServer.Start()
	if err != nil {
		log.Fatalf("Error starting sub server: %v", err)
		return
	}
	// 这行代码创建了一个缓冲为 1 的信号通道 sigCh，用来接收操作系统发出的信号。
	sigCh := make(chan os.Signal, 1)
	// Trap shutdown signals
	// 使用 signal.Notify 函数让 Go 进程监听 SIGHUP 和 SIGTERM 信号。这意味着当操作系统发送这两种信号时，信号会被发送到 sigCh 通道中。
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-sigCh

		switch sig {
		case syscall.SIGHUP:
			// logger.Info("Received SIGHUP signal. Restarting servers...")

		case syscall.SIGINT:
			// 添加处理 SIGINT 时的逻辑
			log.Println("Received SIGINT signal. Shutting down servers. ")
			server.Stop()
			subServer.Stop()
			log.Println("Stopped server")
			// 必须要返回，不返回就会一直拦截相关的信号量
			return
		default:
			server.Stop()
			subServer.Stop()
			log.Println("Shutting down servers.")
			return
		}
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
