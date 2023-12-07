package adapter

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// Define the struct that adapt HTTP request to Beego controller
type BeegoControllerAdapter struct {
	routeInfo  map[string]string
	ctrlName   string
	controller web.ControllerInterface
}

func NewBeegoControllerAdapter(ctrl web.ControllerInterface, info map[string]string, name string) *BeegoControllerAdapter {
	return &BeegoControllerAdapter{
		routeInfo:  info,
		ctrlName:   name,
		controller: ctrl,
	}
}

// HandleRequest adapts the Azure Functions HTTP trigger to Beego controller
func (adapter *BeegoControllerAdapter) HandleRequest(w http.ResponseWriter, r *http.Request, actionName string) {
	// Create a Beego context
	ctx := context.NewContext()
	ctx.Reset(w, r)

	// Create a minimal Beego app instance
	app := web.NewHttpSever()

	// Initialize the Beego controller with the action name and minimal app instance
	adapter.controller.Init(ctx, adapter.ctrlName, actionName, app)

	// Dynamically call the appropriate method based on the action name
	switch actionName {
	case "Get":
		adapter.controller.Get()
	case "Post":
		adapter.controller.Post()
	case "Put":
		adapter.controller.Put()
	case "Delete":
		adapter.controller.Delete()
	case "Prepare":
		adapter.controller.Prepare()
	}
}

// Trigger handler functions for each HTTP method
func (adapter *BeegoControllerAdapter) GetTriggerHandler(w http.ResponseWriter, r *http.Request) {
	adapter.HandleRequest(w, r, "Get")
}

func (adapter *BeegoControllerAdapter) PostTriggerHandler(w http.ResponseWriter, r *http.Request) {
	adapter.HandleRequest(w, r, "Post")
}

func (adapter *BeegoControllerAdapter) PutTriggerHandler(w http.ResponseWriter, r *http.Request) {
	adapter.HandleRequest(w, r, "Put")
}

func (adapter *BeegoControllerAdapter) DeleteTriggerHandler(w http.ResponseWriter, r *http.Request) {
	adapter.HandleRequest(w, r, "Delete")
}

func (adapter *BeegoControllerAdapter) PrepareTriggerHandler(w http.ResponseWriter, r *http.Request) {
	adapter.HandleRequest(w, r, "Prepare")
}
