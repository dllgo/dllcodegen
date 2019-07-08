package codegen

import "html/template"

var handlerTmpl = template.Must(template.New("handler").Parse(handlerTmplCode))

const handlerTmplCode = `
package handler
import (
	"github.com/gin-gonic/gin"
	"{{.PkgName}}/services"
	"github.com/dllgo/dllkit/gins"
)
type {{.Name}}Handler struct {
	{{.Name}}Service      *services.{{.Name}}Service
}
func New{{.Name}}Handler() *{{.Name}}Handler {
	return &{{.Name}}Handler {
        {{.Name}}Service: services.New{{.Name}}Service(),
    }
}
func (this *{{.Name}}Handler) Router(router *gin.Engine) {
	r := router.Group("{{.Name}}")
	r.GET("test", this.test)

}
func (this *{{.Name}}Handler) test(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)

	result.Code = 200
	result.Msg = "成功"
	return 
}
`
