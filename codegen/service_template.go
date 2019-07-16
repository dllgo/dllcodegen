package codegen

import "html/template"

var serviceTmpl = template.Must(template.New("service").Parse(serviceTmplCode))

const serviceTmplCode = `
package services
import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/repositories"
	"github.com/dllgo/dllkit/db/gorm"
	"github.com/dllgo/dllkit/gins"
)
type {{.Name}}Service struct {
	{{.Name}}Repository *repositories.{{.Name}}Repository
}
func New{{.Name}}Service() *{{.Name}}Service {
	return &{{.Name}}Service {
        {{.Name}}Repository: repositories.New{{.Name}}Repository(),
    }
}
func (this *{{.Name}}Service) Get(id int64) *model.{{.Name}} {
	return this.{{.Name}}Repository.Get(gorm.MustDB(), id)
}
func (this *{{.Name}}Service) Count(query map[string]interface{}) int {
	return this.{{.Name}}Repository.Count(gorm.MustDB(),query)
}
func (this *{{.Name}}Service) Query(pageNo int, pageSize int,query map[string]interface{}) (list []model.{{.Name}}, paging *gins.Paging) {
	return this.{{.Name}}Repository.Query(gorm.MustDB(), pageNo ,pageSize ,query)
}
func (this *{{.Name}}Service) Take(where ...interface{}) *model.{{.Name}} {
	return this.{{.Name}}Repository.Take(gorm.MustDB(), where...)
}
func (this *{{.Name}}Service) Create(t *model.{{.Name}}) error {
	return this.{{.Name}}Repository.Create(gorm.MustDB(), t)
}
func (this *{{.Name}}Service) Update(t *model.{{.Name}}) error {
	return this.{{.Name}}Repository.Update(gorm.MustDB(), t)
}
func (this *{{.Name}}Service) UpdateColumn(id int64, columns map[string]interface{}) error {
	return this.{{.Name}}Repository.UpdateColumn(gorm.MustDB(), id, columns)
}
func (this *{{.Name}}Service) Delete(ids ...int64) {
	this.{{.Name}}Repository.Delete(gorm.MustDB(), ids...)
}

`
