package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xianlubird/orchestrator/routers"
)

func main() {
	beego.Run()
}
