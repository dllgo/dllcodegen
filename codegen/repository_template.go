package codegen

import "html/template"

var repositoryTmpl = template.Must(template.New("repository").Parse(repositoryTmplCode))

const repositoryTmplCode = `
package repositories
import (
	"{{.PkgName}}/model"
	"github.com/jinzhu/gorm"
	"github.com/dllgo/dllkit/gins"
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

func (this *{{.Name}}Repository) Count(db *gorm.DB,query map[string]interface{}) int {
	ret := &model.{{.Name}}{}
	var count int
	if err := db.Model(ret).Where(query).Count(&count).Error; err != nil {
		return 0
	}
	return count
}
func (this *{{.Name}}Repository) Query(db *gorm.DB,pageNo int, pageSize int,query map[string]interface{}) (list []model.{{.Name}}, paging *gins.Paging) {
	page := gins.NewPaging()
	page.PageNo = pageNo
	page.PageSize = pageSize
	page.TotalCount = this.Count(db,query)
	
	if page.TotalCount ==0 {
		return nil, nil
	}
	page.TotalPage = page.TotalPages()

	var {{.Name}}s []model.{{.Name}}
	err := db.Where(query).Offset(page.Offset()).Limit(page.PageSize).Order("id desc").Find(&{{.Name}}s).Error
	if err != nil {
		return nil, nil
	}
	return {{.Name}}s, page
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
func (this *{{.Name}}Repository) UpdateColumn(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}
func (this *{{.Name}}Repository) Delete(db *gorm.DB, ids ...int64) {
	db.Where("id in (?)", ids).Delete(&model.{{.Name}}{})
}
`
