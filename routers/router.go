package routers

import (
	"github.com/xianlubird/haha/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
