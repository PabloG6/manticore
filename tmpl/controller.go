 
package tmpl

type ControllerCommandArgs struct {
	ControllerName string;
}
func ControllerTemplate() []byte {
	return []byte(`
package controllers	
type {{.ControllerName | ToUpperCase}}Controller struct {

}

func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Create() {

}


func ({{.ControllerName | ToLowerCase}} *{{.ControllerName | ToUpperCase}}Controller) Update() {

}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Delete() {

}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Show() {

}


func ({{.ControllerName | ToLowerCase}}Controller *{{.ControllerName | ToUpperCase}}Controller) Patch() {

}
	`)
}
	
