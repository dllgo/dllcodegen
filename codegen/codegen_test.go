package codegen

import (
	"fmt"
	"github.com/dllgo/dllcodegen/example/model"
	"testing"
)

func TestCodegen(t *testing.T) {
	fmt.Println("1212")
	Generate("../", "github.com/dllgo/dllcodegen", GetGenerateStruct(&model.Menu{}))
}
