package controller

import (
	"x-ui-scratch/logger"

	"github.com/gin-gonic/gin"
)

type XUIController struct {
	BaseController
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
	logger.Info("TODO: add init router")
}

func (a *XUIController) index(c *gin.Context) {
	html(c, "index.html", "pages.index.title", nil)
}
