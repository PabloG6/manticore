package manticore

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func(app *App) GET(path string, handler gin.HandlerFunc) gin.IRoutes{
	return app.Routes.GET(fmt.Sprintf("%s/:id", path), handler)
}

func(app *App) ALL(path string, handler gin.HandlerFunc) gin.IRoutes {
	return app.Routes.GET(path, handler)
}


func(app *App) POST(path string, handler gin.HandlerFunc) gin.IRoutes{
	return app.Routes.POST(path, handler)
}


func(app *App) PATCH(path string, handler gin.HandlerFunc) gin.IRoutes{
	return app.Routes.PATCH(path, handler)
}

func(app *App) DELETE(path string, handler gin.HandlerFunc) gin.IRoutes {
	return app.Routes.DELETE(path, handler)
}

func(app *App) PUT(path string, handler gin.HandlerFunc) gin.IRoutes {
	return app.Routes.PUT(path, handler)
}