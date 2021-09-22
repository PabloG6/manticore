package tests

import (

	"testing"

	"github.com/PabloG6/manticore"
	"github.com/gin-gonic/gin"
)

func Test_RouteMapping_RoutesFn(testing *testing.T) {
	 r := manticore.New();
	 routerFn := manticore.RoutesFn{
	 	Create: func(c *gin.Context) {
	 	},
	 	Delete: func(c *gin.Context) {
	 	},
	 	Put: func(c *gin.Context) {
	 	},
	 	Patch: func(c *gin.Context) {
	 	},
	 	Get: func(c *gin.Context) {
	 	},
	 }
	 r.Router("user", routerFn)
}

func Create(c *gin.Context) {
	
}