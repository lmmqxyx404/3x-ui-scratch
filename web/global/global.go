package global

var (
	webServer WebServer
	// subServer SubServer
)

type WebServer interface {
	// GetCron() *cron.Cron
	// GetCtx() context.Context
}

func SetWebServer(s WebServer) {
	webServer = s
}
