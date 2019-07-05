package codegen

import "html/template"

var handlerTmpl = template.Must(template.New("handler").Parse(handlerTmplCode))

const handlerTmplCode = `
package handler
import (
	"{{.PkgName}}/services"
)
type {{.Name}}Handler struct {
	{{.Name}}Service      *services.{{.Name}}Service
}
func New{{.Name}}Handler() *{{.Name}}Handler {
	return &{{.Name}}Handler {
        {{.Name}}Service: services.New{{.Name}}Service(),
    }
}
`
