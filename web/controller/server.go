package controller

import (
	"time"
	"x-ui-scratch/logger"
	"x-ui-scratch/web/global"
	"x-ui-scratch/web/service"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	lastGetStatusTime time.Time
	lastStatus        *service.Status
	serverService     service.ServerService

	BaseController
}

func NewServerController(g *gin.RouterGroup) *ServerController {
	a := &ServerController{
		lastGetStatusTime: time.Now(),
	}
	a.initRouter(g)
	a.startTask()
	return a
}

func (a *ServerController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/server")
	logger.Info("TODO: initRouter")
	g.Use(a.checkLogin)
	g.POST("/status", a.status)
	/* g.POST("/getXrayVersion", a.getXrayVersion)
	g.POST("/stopXrayService", a.stopXrayService)
	g.POST("/restartXrayService", a.restartXrayService)
	g.POST("/installXray/:version", a.installXray)
	g.POST("/logs/:count", a.getLogs)
	g.POST("/getConfigJson", a.getConfigJson)
	g.GET("/getDb", a.getDb)
	g.POST("/importDB", a.importDB)
	g.POST("/getNewX25519Cert", a.getNewX25519Cert) */
}

func (a *ServerController) startTask() {
	webServer := global.GetWebServer()
	c := webServer.GetCron()
	c.AddFunc("@every 10s", func() {
		now := time.Now()
		if now.Sub(a.lastGetStatusTime) > time.Minute*3 {
			return
		}
		a.refreshStatus()
	})
}

func (a *ServerController) refreshStatus() {
	a.lastStatus = a.serverService.GetStatus(a.lastStatus)
}

func (a *ServerController) status(c *gin.Context) {
	a.lastGetStatusTime = time.Now()

	jsonObj(c, a.lastStatus, nil)
}
