package common

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *App) ShowRoutes() interface{} {
	return app.routes
}

func (app *App) AddRoute(method, path string, handler gin.HandlerFunc) {
	app.routes = append(
		app.routes,
		Route{strings.ToUpper(method), path, handler},
	)
}

func (app *App) OptionsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, email, password, username")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "HEAD, PUT, GET, POST, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.Status(204)
	}
}

func (app *App) Serve() error {

	filter := map[string]bool{}
	for _, route := range app.routes {
		if !filter[route.Path] {
			app.Gin.OPTIONS(route.Path, app.OptionsHandler)
			filter[route.Path] = true
		}
		app.Gin.Handle(route.Method, route.Path, app.OptionsHandler, route.Handler)
	}

	s := &http.Server{
		Addr:           "0.0.0.0:" + os.Getenv("PORT"),
		Handler:        app.Gin,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	log.Println("SERVER SHUTTING DOWN...")
	return err
}
