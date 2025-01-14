package controller

import (
	"x-ui-scratch/web/service"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	settingService service.SettingService
	/* userService    service.UserService
	panelService   service.PanelService */
}

func NewSettingController(g *gin.RouterGroup) *SettingController {
	a := &SettingController{}
	a.initRouter(g)
	return a
}

func (a *SettingController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/setting")

	g.POST("/defaultSettings", a.getDefaultSettings)

}

func (a *SettingController) getDefaultSettings(c *gin.Context) {
	result, err := a.settingService.GetDefaultSettings(c.Request.Host)
	if err != nil {
		jsonMsg(c, I18nWeb(c, "pages.settings.toasts.getSettings"), err)
		return
	}
	jsonObj(c, result, nil)
}
