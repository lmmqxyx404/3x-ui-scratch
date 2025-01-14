package controller

import (
	"x-ui-scratch/logger"

	"github.com/gin-gonic/gin"
)

type XUIController struct {
	BaseController

	settingController *SettingController
}

func NewXUIController(g *gin.RouterGroup) *XUIController {
	a := &XUIController{}
	a.initRouter(g)
	return a
}

func (a *XUIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/panel")
	g.Use(a.checkLogin)
	g.GET("/", a.index)
	// note: vital router
	g.GET("/inbounds", a.inbounds)
	// g.GET("/settings", a.settings)

	a.settingController = NewSettingController(g)

	logger.Info("TODO: add init router")

}

func (a *XUIController) index(c *gin.Context) {
	html(c, "index.html", "pages.index.title", nil)
}

func (a *XUIController) inbounds(c *gin.Context) {
	html(c, "inbounds.html", "pages.inbounds.title", nil)
}
