package manticore

import "github.com/gin-gonic/gin"

type RoutesFn struct {
	Create gin.HandlerFunc
	Delete gin.HandlerFunc
	Put gin.HandlerFunc
	Patch gin.HandlerFunc
	Get gin.HandlerFunc
}

/**
quickly bootstrap a list of controller functions with the routesFn struct.
*/
func (app *App) Router(path string, routes RoutesFn) {
	app.GET(path, routes.Get)
	app.POST(path, routes.Create)
	app.PATCH(path, routes.Patch)
	app.DELETE(path, routes.Delete)
	app.PUT(path, routes.Put)
}