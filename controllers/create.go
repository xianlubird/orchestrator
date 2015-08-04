package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xianlubird/orchestrator/utils"
	"log"
)

type ForwardController struct {
	beego.Controller
}

func (this *ForwardController) Get() {
	this.TplNames = "createContainer.html"
}

func (this *ForwardController) Post() {
	codePath := this.GetString("code_path")
	log.Println("codePath", codePath)
	containerID := utils.CreateContainer()
	log.Println("containerID", containerID)
	utils.StartContainer(containerID)
}