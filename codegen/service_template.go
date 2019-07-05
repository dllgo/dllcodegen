package codegen

import "html/template"

var serviceTmpl = template.Must(template.New("service").Parse(serviceTmplCode))

const serviceTmplCode = `
package services
import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/repositories"
	"github.com/dllgo/dllkit/db/gorm"
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
func (this *{{.Name}}Service) GetInIds(Ids []int64) []model.{{.Name}} {
	return this.{{.Name}}Repository.GetInIds(gorm.MustDB(), Ids)
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
func (this *{{.Name}}Service) Updates(id int64, columns map[string]interface{}) error {
	return this.{{.Name}}Repository.Updates(gorm.MustDB(), id, columns)
}
func (this *{{.Name}}Service) UpdateColumn(id int64, name string, value interface{}) error {
	return this.{{.Name}}Repository.UpdateColumn(gorm.MustDB(), id, name, value)
}
func (this *{{.Name}}Service) Delete(id int64) {
	this.{{.Name}}Repository.Delete(gorm.MustDB(), id)
}
func (this *{{.Name}}Service) DeleteInIds(Ids []int64) {
	this.{{.Name}}Repository.DeleteInIds(gorm.MustDB(), Ids)
}
`
