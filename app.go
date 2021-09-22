package manticore

import (
	"os"

	"github.com/gin-gonic/gin"
)
type App struct {
	Routes *gin.Engine

}


func New() *App {

	app := &App{
		Routes: gin.New(),
		
	}
	return app;
}

func Default() *App {
	app := &App{
		Routes: gin.Default(),
	}
	return app;
}


func (app *App) Start(port ...string) {
	addr := resolveAddress(port)
	app.Routes.Run(addr)
}


func resolveAddress(addr[] string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			return ":" + port
		}
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}


