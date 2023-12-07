package main

import (
	azfuncAdapter "beegoserverless/serverlessAdapter"

	"test/controllers"
)

func main() {
	shortCtrl := &controllers.ShortController{routeInfo: make(map[string]string)}
	shortCtrl.routeInfo["/api/shorten"] = "Get"
	expandtrl := &controllers.ExpandController{routeInfo: make(map[string]string)}
	expandtrl.routeInfo["/api/expand"] = "Get"

	azfuncAdapter.HttpFuncsHandler(shortCtrl.routeInfo, shortCtrl, "shorturl")
	azfuncAdapter.HttpFuncsHandler(shortCtrl.routeInfo, shortCtrl, "expandurl")
	azfuncAdapter.PortHandler(":8080")
}
