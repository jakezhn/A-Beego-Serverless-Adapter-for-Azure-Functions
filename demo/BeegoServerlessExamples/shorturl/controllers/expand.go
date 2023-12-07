package controllers

import (
	"github.com/astaxie/beego"
)

type ExpandController struct {
	beego.Controller
	routeInfo map[string]string
}

func (this *ExpandController) Get() {
	var result ShortResult
	shorturl := this.Input().Get("shorturl")
	result.UrlShort = shorturl
	if urlcache.IsExist(shorturl) {
		result.UrlLong = urlcache.Get(shorturl).(string)
	} else {
		result.UrlLong = ""
	}
	this.Data["json"] = result
	this.ServeJSON()
}
