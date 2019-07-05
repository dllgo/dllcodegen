package codegen

import "html/template"

var repositoryTmpl = template.Must(template.New("repository").Parse(repositoryTmplCode))

const repositoryTmplCode = `
package repositories
import (
	"{{.PkgName}}/model"
	"github.com/jinzhu/gorm"
)
type {{.Name}}Repository struct {
}
func New{{.Name}}Repository() *{{.Name}}Repository {
	return &{{.Name}}Repository{}
}
func (this *{{.Name}}Repository) Get(db *gorm.DB, id int64) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
func (this *{{.Name}}Repository) GetInIds(db *gorm.DB,Ids []int64) []model.{{.Name}} {
	if len(Ids) == 0 {
		return nil
	}
	var {{.Name}}s []model.{{.Name}}
	if err := db.Where("id in (?)", Ids).Find(&{{.Name}}s).Error; err != nil {
		return nil
	}
	return {{.Name}}s
}
func (this *{{.Name}}Repository) Take(db *gorm.DB, where ...interface{}) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (this *{{.Name}}Repository) Create(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Create(t).Error
	return
}
func (this *{{.Name}}Repository) Update(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Save(t).Error
	return
}
func (this *{{.Name}}Repository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}
func (this *{{.Name}}Repository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
func (this *{{.Name}}Repository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.{{.Name}}{}).Delete("id", id)
}
func (this *{{.Name}}Repository) DeleteInIds(db *gorm.DB,Ids []int64) {
	db.Where("id in (?)", Ids).Delete(&model.{{.Name}}{})
}
`
