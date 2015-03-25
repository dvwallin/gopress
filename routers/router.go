package routers

import (
	"github.com/dvwallin/gopress/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
