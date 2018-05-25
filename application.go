package bibo

import (
	"flag"
	"net/http"
	"log"
	"github.com/ketu/bibo/routing"
	"time"
	"github.com/ketu/bibo/middleware"
)

var (
	addr = flag.String("addr", ":8999", "TCP address to listen to")
)

type Application struct {
	configuration Configuration
	routing       *routing.Routing
	middleware   map[string]middleware.Middleware
}

func Default() *Application {
	configuration := DefaultConfiguration()
	routes := routing.NewRouting()
	return &Application{
		routing:       routes,
		configuration: configuration,
	}
}

func (app Application) ListenAndServe() {
	flag.Parse()
	//http.HandleFunc("/", transferHandleFunc)
	//handlers = app.router.Dump()
	//routes := app.routing.Dump()

	srv := &http.Server{
		Handler:      app.routing,
		Addr:         *addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
