package adapter

import (
	"log"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

// HttpFuncsHandler sets up HTTP handlers based on the provided route information (url and action)
func HttpFuncsHandler(routeInfo map[string]string, beegoCtrl web.ControllerInterface, ctrlName string) {
	adapter := NewBeegoControllerAdapter(beegoCtrl, routeInfo, ctrlName)

	for path, method := range routeInfo {
		// Register the route with Beego's router
		if err := web.Router(path, beegoCtrl); err != nil {
			log.Printf("Error setting up Beego router for path %s: %v", path, err)
		}

		// Register the route with the HTTP server based on the action
		switch method {
		case "Get":
			http.HandleFunc(path, adapter.GetTriggerHandler)
		case "Post":
			http.HandleFunc(path, adapter.PostTriggerHandler)
		case "Put":
			http.HandleFunc(path, adapter.PutTriggerHandler)
		case "Delete":
			http.HandleFunc(path, adapter.DeleteTriggerHandler)
		case "Prepare":
			http.HandleFunc(path, adapter.PrepareTriggerHandler)
		default:
			log.Printf("Unsupported method %s for path %s", method, path)
		}

		log.Printf("Route configured: %s %s", method, path)
	}
}

// As in cloud environments like Azure Functions, ports may be dynamically allocated
// PortHandler starts the HTTP server on the specified address
// It checks for a custom port setting and logs the server start
func PortHandler(listenAddr string) {
	// Check if a custom port is set in the environment variables
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
		log.Printf("Custom handler port detected, using port: %s", val)
	}

	// Log a message indicating the server is about to start
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)

	// Start the HTTP server and listen on the specified address
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server on %s: %v", listenAddr, err)
	}
}
