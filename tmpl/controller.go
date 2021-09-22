 
package tmpl

type ControllerCommandArgs struct {
	ControllerName string;
}
func ControllerTemplate() []byte {
	return []byte(`
package controllers	
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type {{.ControllerName | ToUpperCase}}Controller struct {

}

func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})

}


func ({{.ControllerName | ToLowerCase}} *{{.ControllerName | ToUpperCase}}Controller) Update(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})

}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})


}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})

}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Patch(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})


}
	`)
}
	
