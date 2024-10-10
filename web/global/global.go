package global

var (
	webServer WebServer
	subServer SubServer
)

type WebServer interface {
	// GetCron() *cron.Cron
	// GetCtx() context.Context
}

type SubServer interface {
	// GetCtx() context.Context
}

func SetWebServer(s WebServer) {
	webServer = s
}

func SetSubServer(s SubServer) {
	subServer = s
}
