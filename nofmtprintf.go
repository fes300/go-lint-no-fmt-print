package linters

import (
	"go/ast"
	"go/types"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("nofmtprintf", New)
}

type MySettings struct{}

type Plugin struct {
	settings MySettings
}

func New(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	s, err := register.DecodeSettings[MySettings](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{settings: s}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "nofmtprintf",
			Doc:  "disallows printing with the fmt package",
			Run:  f.run,
		},
	}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func (f *Plugin) run(pass *analysis.Pass) (interface{}, error) {
	var fmtPrintfType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Name() == "fmt" {
			fmtPrintfType = pkg.Scope().Lookup("Printf").Type()
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// examine all function calls
			callExpr, isCallExpr := node.(*ast.CallExpr)
			if !isCallExpr {
				return true
			}

			callExprFunType := pass.TypesInfo.TypeOf(callExpr.Fun)
			if callExprFunType == fmtPrintfType {
				pass.Report(analysis.Diagnostic{
					Pos:     node.Pos(),
					End:     node.End(),
					Message: "Don't use fmt.Printf",
				})
			}

			return true
		})
	}

	return nil, nil
}
