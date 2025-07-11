package wpsfuncprefix

import (
	"github.com/golangci/plugin-module-register/register"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func init() {
	register.Plugin("wpsfuncprefix", New)
}

type PluginWPSFuncPrefix struct{}

func (f *PluginWPSFuncPrefix) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "wps_func_prefix",
			Doc:  "check that function names start with WPS_",
			Run:  run,
		},
	}, nil
}

func (f *PluginWPSFuncPrefix) GetLoadMode() string {
	return register.LoadModeSyntax
}

func New(settings any) (register.LinterPlugin, error) {

	return &PluginWPSFuncPrefix{}, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		// 跳过测试文件
		if strings.HasSuffix(pass.Fset.File(file.Pos()).Name(), "_test.go") {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			// 跳过 init 函数和 main 函数
			if fn.Name.Name == "init" || fn.Name.Name == "main" {
				return true
			}

			// 检查函数名是否以 WPS_ 开头
			if !strings.HasPrefix(fn.Name.Name, "WPS_") {
				pass.Reportf(
					fn.Name.Pos(),
					"function %q should start with WPS_",
					fn.Name.Name,
				)
			}
			return true
		})
	}
	return nil, nil
}
