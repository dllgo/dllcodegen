package codegen

import (
	"fmt"
	"github.com/dllgo/dllcodegen/model"

	//"github.com/dllgo/dllcodegen/services"
	"testing"
)

func TestCodegen(t *testing.T) {
	fmt.Println("1212")
	Generate("../","github.com/dllgo/dllcodegen",GetGenerateStruct(&model.Words{}))

	//us := services.NewWordsService()
	//fmt.Println(us.Get(1))
}
