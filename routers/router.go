package routers

import (
	"github.com/xianlubird/orchestrator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/create", &controllers.ForwardController{})
}
