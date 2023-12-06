package main

import (
	azfuncAdapter "beegoserverless/serverlessAdapter"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// Controller is controller for BeegoApp
type MyController1 struct {
	web.Controller
	routeInfo map[string]string
}

type MyController2 struct {
	web.Controller
	routeInfo map[string]string
}

// Define the Get method of Beego controller, response to HTTP GET request
func (ctrl *MyController1) Get() {
	// Print "Hello from Beego Controller" as response body
	ctrl.Ctx.WriteString("Hello from Beego Controller!")
}

func (ctrl *MyController2) Post() {
	// Beego automatically checks the HTTP method, so no need to check it here

	if ctrl.Ctx.Request.Method != "POST" {
		ctrl.Ctx.Output.SetStatus(405) // HTTP 405 Method Not Allowed
		ctrl.Ctx.Output.Body([]byte("Only POST method is accepted"))
		return
	}
	// Parse JSON body into a map
	var data map[string]interface{}
	body, err := io.ReadAll(ctrl.Ctx.Request.Body)
	if err != nil {
		ctrl.Ctx.Output.SetStatus(500)
		ctrl.Ctx.Output.Body([]byte("Error reading request body"))
		return
	}

	defer ctrl.Ctx.Request.Body.Close()
	if err := json.Unmarshal(body, &data); err != nil {
		ctrl.Ctx.Output.SetStatus(400)
		ctrl.Ctx.Output.Body([]byte("Error parsing JSON body"))
		return
	}

	// Iterate over the map and output the values
	for _, value := range data {
		if value == "How about you?" {
			ctrl.Ctx.WriteString(fmt.Sprintf("%v", "Whatsup from Azure Functions!"))
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from azfunction!")
}

func main() {
	// Create controllers and set up routers' indo (path and action)
	Ctrl1 := &MyController1{routeInfo: make(map[string]string)}
	Ctrl1.routeInfo["/api/hello"] = "Get"

	Ctrl2 := &MyController2{routeInfo: make(map[string]string)}
	Ctrl2.routeInfo["/api/whatsup"] = "Post"

	// Register routers and http functions
	azfuncAdapter.HttpFuncsHandler(Ctrl1.routeInfo, Ctrl1, "MyController1")
	azfuncAdapter.HttpFuncsHandler(Ctrl2.routeInfo, Ctrl2, "MyController2")

	//Setting up default port listening
	azfuncAdapter.PortHandler(":8080")
}
