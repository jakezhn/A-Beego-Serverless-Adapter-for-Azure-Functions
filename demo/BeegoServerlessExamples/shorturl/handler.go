package main

import (
	azfuncAdapter "beegoserverless/serverlessAdapter"

	"github.com/beego/samples/shorturl/controllers"
)

func main() {
	mainCtrl := &controllers.MainController{routeInfo: make(map[string]string)}
	mainCtrl.routeInfo["/"] = "Get"
	shortCtrl := &controllers.ShortController{routeInfo: make(map[string]string)}
	shortCtrl.routeInfo["/v1/shorten"] = "Get"
	expandtrl := &controllers.ExpandController{routeInfo: make(map[string]string)}
	expandtrl.routeInfo["/v1/expand"] = "Get"

	azfuncAdapter.HttpFuncsHandler(mainCtrl.routeInfo, mainCtrl, "main")
	azfuncAdapter.HttpFuncsHandler(shortCtrl.routeInfo, shortCtrl, "shorturl")
	azfuncAdapter.HttpFuncsHandler(shortCtrl.routeInfo, shortCtrl, "expandurl")
	azfuncAdapter.PortHandler(":8080")
}
