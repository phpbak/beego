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
    beego.AutoRouter(&controllers.ConfController{})
    beego.AutoRouter(&controllers.LogController{})
    beego.AutoRouter(&controllers.CacheController{})
    beego.AutoRouter(&controllers.LoginController{})
    beego.AutoRouter(&controllers.LayoutController{})
    beego.AutoRouter(&controllers.DbController{})  

    beego.AutoRouter(&controllers.ApiController{})
    beego.AutoRouter(&controllers.ApiJsonController{})
    beego.AutoRouter(&controllers.ApiXMLController{})
    beego.AutoRouter(&controllers.ApiJsonpController{})
    beego.AutoRouter(&controllers.ApiDictionaryController{})
    beego.AutoRouter(&controllers.ApiParamsController{})
}
