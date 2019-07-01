package routers

import (
	"PPGo_amaze/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 默认登录
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/stockindex399", &controllers.Stockindex{}, "*:Index399")
	beego.Router("/stockindex300", &controllers.Stockindex{}, "*:Index300")
	beego.Router("/stockindex", &controllers.Stockindex{}, "*:Index")

}
