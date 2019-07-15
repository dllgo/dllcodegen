package codegen

import "html/template"

var handlerTmpl = template.Must(template.New("handler").Parse(handlerTmplCode))

const handlerTmplCode = `
package handler
import (
	"github.com/gin-gonic/gin"
	"{{.PkgName}}/services"
	"{{.PkgName}}/model"
	"github.com/dllgo/dllkit/gins"
	"github.com/dllgo/dllkit/os/conv"
	"github.com/dllgo/dllkit/validation"
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
	r.GET("/lists", this.get{{.Name}}s)
	r.GET("/one/:id", this.get{{.Name}})
	r.POST("/add", this.add{{.Name}})
	r.POST("/edit/:id", this.edit{{.Name}})
	r.POST("/delete/:id", this.delete{{.Name}})
	r.POST("/batchdel", this.delete{{.Name}}s)
}
func (this *{{.Name}}Handler) get{{.Name}}s(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)

	result.Code = 200
	result.Msg = "成功"
	return
}
func (this *{{.Name}}Handler) get{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	id := conv.Int64(ctx.Param("id"))
	valid := validation.Validation{}
	if v := valid.Min(id, 1, "id"); !v.Ok {
		result.Code = 300
		result.Msg = "ID必须大于0"
		return
	}
	result.Code = 200
	result.Msg = "成功"
	result.Data = this.{{.Name}}Service.Get(id)
	return
}
func (this *{{.Name}}Handler) add{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	var {{.Name}}Info model.{{.Name}}
	err := ctx.BindJSON(&{{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "无效的参数"
		return
	}
	valid := validation.Validation{}
	valid.MaxSize({{.Name}}Info.Name, 20, "name").Message("最长为20字符")
	if valid.HasErrors() {
		result.Msg = "无效的参数"
		result.Data = valid.Errors
		return
	}
	m{{.Name}} := this.{{.Name}}Service.Take({{.Name}}Info)
	if m{{.Name}}.Id > 0 {
		result.Code = 300
		result.Msg = "记录已存在"
		return
	}

	err := this.{{.Name}}Service.Create({{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "新增失败"
		return
	}
	result.Code = 200
	result.Msg = "新增成功"
	return
}
func (this *{{.Name}}Handler) edit{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	id := conv.Int64(ctx.Param("id"))
	valid := validation.Validation{}
	if v := valid.Min(id, 1, "id"); !v.Ok {
		result.Code = 300
		result.Msg = "ID必须大于0"
		return
	}

	var {{.Name}}Info model.{{.Name}}
	err := ctx.BindJSON(&{{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "无效的参数"
		return
	}
	{{.Name}}Info.Id = id
	err := this.{{.Name}}Service.Update({{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "更新失败"
		return
	}
	result.Code = 200
	result.Msg = "更新成功"
	return
}
func (this *{{.Name}}Handler) delete{{.Name}}s(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)

	result.Code = 200
	result.Msg = "成功"
	return
}
func (this *{{.Name}}Handler) delete{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	id := conv.Int64(ctx.Param("id"))
	valid := validation.Validation{}
	if v := valid.Min(id, 1, "id"); !v.Ok {
		result.Code = 300
		result.Msg = "ID必须大于0"
		return
	}
	this.{{.Name}}Service.Delete(id)
	result.Code = 200
	result.Msg = "删除成功"
	return
}
`
