package codegen

import (
	"bytes"
	"fmt"
	"github.com/dllgo/dllcodegen/utils"
	"os"
	"path"
	"reflect"
)

type GenerateStruct struct {
	Name   string
	Fields []GenerateField
}

type GenerateField struct {
	CamelName   string
	NativeField reflect.StructField
}

type InputData struct {
	PkgName   string
	Name      string
	KebabName string // FuckShit -> fuck-shit
	Fields    []GenerateField
}

func Generate(baseDir, pkgName string, models ...GenerateStruct) {
	for _, model := range models {
		if err := generateRepository(baseDir, pkgName, model); err != nil {
			fmt.Printf(err.Error())
		}
		if err := generateService(baseDir, pkgName, model); err != nil {
			fmt.Printf(err.Error())
		}
		if err := generateHandler(baseDir, pkgName, model); err != nil {
			fmt.Printf(err.Error())
		}
	}
}

func GetGenerateStruct(s interface{}) GenerateStruct {
	structName := utils.StructName(s)
	structFields := utils.StructFields(s)

	var fields []GenerateField
	for _, f := range structFields {
		if f.Anonymous {
			continue
		}
		fields = append(fields, GenerateField{
			CamelName:   utils.ToLowerCamel(f.Name),
			NativeField: f,
		})
	}

	return GenerateStruct{
		Name:   structName,
		Fields: fields,
	}
}

func generateRepository(baseDir, pkgName string, s GenerateStruct) error {
	var b bytes.Buffer
	repositoryTmpl.Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		KebabName: utils.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	c := b.String()

	path, err := getFilePath(baseDir, "/repositories/"+utils.ToSnake(s.Name+"_repository.go"))
	if err != nil {
		return err
	}
	return writeFile(path, c)
}

func generateService(baseDir, pkgName string, s GenerateStruct) error {
	var b bytes.Buffer
	serviceTmpl.Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		KebabName: utils.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	c := b.String()

	path, err := getFilePath(baseDir, "/services/"+utils.ToSnake(s.Name+"_service.go"))
	if err != nil {
		return err
	}
	return writeFile(path, c)
}
func generateHandler(baseDir, pkgName string, s GenerateStruct) error {
	var b bytes.Buffer
	handlerTmpl.Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		KebabName: utils.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	c := b.String()

	path, err := getFilePath(baseDir, "/handler/"+utils.ToSnake(s.Name+"_handler.go"))
	if err != nil {
		return err
	}
	return writeFile(path, c)
}
func getFilePath(baseDir, sub string) (filepath string, err error) {
	filepath = path.Join(baseDir, sub)
	base := path.Dir(filepath)
	err = os.MkdirAll(base, os.ModePerm)
	return
}

func writeFile(filepath string, content string) error {
	exists, err := utils.PathExists(filepath)
	if err != nil {
		return err
	}
	if exists {
		fmt.Println("文件已经存在...", filepath)
		filepath = filepath + ".temp"
	}
	return utils.WriteString(filepath, content, true)
}
