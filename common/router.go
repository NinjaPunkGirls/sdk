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
	route := Route{strings.ToUpper(method), path, handler}
	app.routes = append(app.routes, route)
	log.Println("adding route:", route.Method, route.Path)
	app.Gin.Handle(route.Method, route.Path, app.OptionsHandler, route.Handler)
}

func (app *App) OptionsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, email, password, username")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "HEAD, PUT, GET, POST, PATCH, DELETE")
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
	}
	port := os.Getenv("PORT")
	s := &http.Server{
		Addr:           "0.0.0.0:" + port,
		Handler:        app.Gin,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("starting server on port:", port)
	err := s.ListenAndServe()
	log.Println("SERVER SHUTTING DOWN...")
	return err
}
