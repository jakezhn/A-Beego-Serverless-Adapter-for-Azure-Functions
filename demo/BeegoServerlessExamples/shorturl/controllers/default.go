package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	routeInfo map[string]string
}

func (this *MainController) Get() {
	this.Ctx.Output.Body([]byte("shorturl"))
}
