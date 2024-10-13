package controller

import (
	"html/template"
	"net/http"
	"x-ui-scratch/logger"
	"x-ui-scratch/web/service"
	"x-ui-scratch/web/session"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	userService    service.UserService
	settingService service.SettingService
}

type LoginForm struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	LoginSecret string `json:"loginSecret" form:"loginSecret"`
}

func NewIndexController(g *gin.RouterGroup) *IndexController {
	a := &IndexController{}
	a.initRouter(g)
	return a
}

func (a *IndexController) initRouter(g *gin.RouterGroup) {
	g.GET("/", a.index)
	logger.Info("TODO: add more routes")
	g.POST("/login", a.login)
	// g.GET("/logout", a.logout)
	// g.POST("/getSecretStatus", a.getSecretStatus)
}

func (a *IndexController) index(c *gin.Context) {
	if session.IsLogin(c) {
		c.Redirect(http.StatusTemporaryRedirect, "panel/")
		logger.Info("redirect to panel/")
		return
	}
	html(c, "login.html", "pages.login.title", nil)
}

func (a *IndexController) login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		pureJsonMsg(c, http.StatusOK, false, I18nWeb(c, "pages.login.toasts.invalidFormData"))
		return
	}
	if form.Username == "" {
		pureJsonMsg(c, http.StatusOK, false, I18nWeb(c, "pages.login.toasts.emptyUsername"))
		return
	}
	if form.Password == "" {
		pureJsonMsg(c, http.StatusOK, false, I18nWeb(c, "pages.login.toasts.emptyPassword"))
		return
	}

	user := a.userService.CheckUser(form.Username, form.Password, form.LoginSecret)
	// timeStr := time.Now().Format("2006-01-02 15:04:05")
	safeUser := template.HTMLEscapeString(form.Username)
	safePass := template.HTMLEscapeString(form.Password)
	safeSecret := template.HTMLEscapeString(form.LoginSecret)
	if user == nil {
		logger.Warningf("wrong username or password or secret: \"%s\" \"%s\" \"%s\"", safeUser, safePass, safeSecret)
		// TODO
		// a.tgbot.UserLoginNotify(safeUser, safePass, getRemoteIp(c), timeStr, 0)
		pureJsonMsg(c, http.StatusOK, false, I18nWeb(c, "pages.login.toasts.wrongUsernameOrPassword"))
		return
	} else {
		logger.Infof("%s logged in successfully, Ip Address: %s\n", safeUser, getRemoteIp(c))
		// TODO
		// a.tgbot.UserLoginNotify(safeUser, ``, getRemoteIp(c), timeStr, 1)
	}

	sessionMaxAge, err := a.settingService.GetSessionMaxAge()
	if err != nil {
		//  TODO
		logger.Info("Unable to get session's max age from DB")
	}

	err = session.SetMaxAge(c, sessionMaxAge*60)
	if err != nil {
		// TODO
		logger.Info("Unable to set session's max age")
	}

	err = session.SetLoginUser(c, user)
	logger.Infof("%s logged in successfully", user.Username)
	jsonMsg(c, I18nWeb(c, "pages.login.toasts.successLogin"), err)
}
