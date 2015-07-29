package routers

import (
	"github.com/astaxie/beego"
	"github.com/xianlubird/orchestrator/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
