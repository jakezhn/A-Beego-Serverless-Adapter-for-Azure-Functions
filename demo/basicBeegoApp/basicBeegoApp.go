package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/beego/beego/v2/server/web"
)

type MyController1 struct {
	web.Controller
}

type MyController2 struct {
	web.Controller
}

func (ctrl *MyController1) Get() {
	// Print "Hello from Beego Controller" as response body
	ctrl.Ctx.WriteString("Hello from Beego Controller!")
}

func (ctrl *MyController2) Post() {
	// Beego automatically checks the HTTP method, so no need to check it here

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
			ctrl.Ctx.WriteString(fmt.Sprintf("%v", "Whatsup from another Beego controller!"))
		}
	}
}

func main() {
	// Create controllers and register routers
	web.Router("/api/hello", &MyController1{})
	web.Router("/api/whatsup", &MyController2{})
	// Run beego server
	web.Run()
}
