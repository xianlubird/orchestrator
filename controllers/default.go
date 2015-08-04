package controllers

import (
	"github.com/astaxie/beego"
	"github.com/samalba/dockerclient"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	
	// Init the client
    docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	
    // Get only running containers
    containers, err := docker.ListContainers(false, false, "")
    if err != nil {
        log.Fatal(err)
    }
    for _, c := range containers {
        log.Println(c.Id, c.Names)
    }
	
	c.TplNames = "index.tpl"
}
