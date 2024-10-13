package controller

import (
	"net/http"
	"x-ui-scratch/logger"
	"x-ui-scratch/web/session"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func NewIndexController(g *gin.RouterGroup) *IndexController {
	a := &IndexController{}
	a.initRouter(g)
	return a
}

func (a *IndexController) initRouter(g *gin.RouterGroup) {
	g.GET("/", a.index)
	logger.Info("TODO: add more routes")
	/* g.POST("/login", a.login)
	g.GET("/logout", a.logout)
	g.POST("/getSecretStatus", a.getSecretStatus) */
}

func (a *IndexController) index(c *gin.Context) {
	if session.IsLogin(c) {
		c.Redirect(http.StatusTemporaryRedirect, "panel/")
		logger.Info("redirect to panel/")
		return
	}
	html(c, "login.html", "pages.login.title", nil)
}
