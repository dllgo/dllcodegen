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
	r.GET("/query/:id", this.get{{.Name}})
	r.POST("/add", this.add{{.Name}})
	r.POST("/edit/:id", this.edit{{.Name}})
	r.POST("/delete", this.delete{{.Name}})
}
func (this *{{.Name}}Handler) get{{.Name}}s(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)

	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")
    var m{{.Name}}Maps map[string]interface{}
	m{{.Name}}s,paging := this.{{.Name}}Service.Query(conv.Int(page), conv.Int(limit),m{{.Name}}Maps)
	if m{{.Name}}s == nil {
		result.Code = 200
		result.Msg = "暂无数据"
		return 
	}
	result.Code = 200
	result.Msg = "成功"
	data := map[string]interface{}{}
	data["lists"] = m{{.Name}}s
	data["page"] = paging
	result.Data = data
	return
}
func (this *{{.Name}}Handler) get{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	mid := conv.Int64(ctx.Param("id"))
	valid := validation.Validation{}
	if v := valid.Min(mid, 1, "id"); !v.Ok {
		result.Code = 300
		result.Msg = "ID必须大于0"
		return
	}
	m{{.Name}} := this.{{.Name}}Service.Get(mid)
	if m{{.Name}} == nil{
		result.Code = 200
		result.Msg = "数据不存在"
		return
	}
	result.Code = 200
	result.Msg = "成功"
	result.Data = m{{.Name}}
	return
}
func (this *{{.Name}}Handler) add{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	var m{{.Name}}Info model.{{.Name}}
	err := ctx.BindJSON(&m{{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "无效的参数"
		return
	}
	valid := validation.Validation{}
	valid.MaxSize(m{{.Name}}Info.Name, 20, "name").Message("最长为20字符")
	if valid.HasErrors() {
		result.Msg = "无效的参数"
		result.Data = valid.Errors
		return
	}
	m{{.Name}} := this.{{.Name}}Service.Take(m{{.Name}}Info)
	if m{{.Name}} != nil && m{{.Name}}.Id > 0 {
		result.Code = 300
		result.Msg = "记录已存在"
		return
	}

	err = this.{{.Name}}Service.Create(&m{{.Name}}Info)
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
	mid := conv.Int64(ctx.Param("id"))
	valid := validation.Validation{}
	if v := valid.Min(mid, 1, "id"); !v.Ok {
		result.Code = 300
		result.Msg = "ID必须大于0"
		return
	}

	var m{{.Name}}Info model.{{.Name}}
	err := ctx.BindJSON(&m{{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "无效的参数"
		return
	}
	m{{.Name}}Info.Id = mid
	err = this.{{.Name}}Service.Update(&m{{.Name}}Info)
	if err != nil {
		result.Code = 300
		result.Msg = "更新失败"
		return
	}
	result.Code = 200
	result.Msg = "更新成功"
	return
}
func (this *{{.Name}}Handler) delete{{.Name}}(ctx *gin.Context) {
	result := gins.NewResponse()
	defer result.JSON(ctx)
	var params map[string]interface{}
	err := ctx.BindJSON(&params)
	if err != nil {
		result.Code = 300
		result.Msg = "无效的参数"
		return
	}
	ids :=conv.String2Int64(params["ids"])
    this.{{.Name}}Service.Delete(ids...)
	result.Code = 200
	result.Msg = "成功"
	return
}
`
