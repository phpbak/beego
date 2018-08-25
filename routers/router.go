package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.FileController{})
    beego.AutoRouter(&controllers.SessionController{})
    beego.AutoRouter(&controllers.JsonController{})
    beego.AutoRouter(&controllers.ConfController{})
    beego.AutoRouter(&controllers.LogController{})
    beego.AutoRouter(&controllers.CacheController{})
    beego.AutoRouter(&controllers.LoginController{})
}
