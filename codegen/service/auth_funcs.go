package service

import (
	"fmt"
	"os"
	"path"
	"strings"

	"goa.design/goa/codegen"
	"goa.design/goa/design"
)

// AuthFuncFiles returns the files that contain a dummy implementation of the
// authorization functions needed to instantiate the service endpoints.
func AuthFuncFiles(genpkg string, root *design.RootExpr) []*codegen.File {
	var result []*codegen.File
	apiPkg := strings.ToLower(codegen.Goify(root.API.Name, false))
	rootPath := "."
	for _, s := range root.Services {
		svc := Services.Get(s.Name)
		if len(svc.Schemes) > 0 {
			var sections []*codegen.SectionTemplate
			specs := []*codegen.ImportSpec{
				{Path: "context"},
				{Path: "fmt"},
				{Path: "goa.design/goa", Name: "goa"},
				{Path: "goa.design/goa/security"},
				{Path: rootPath, Name: apiPkg},
			}
			pkgName := Services.Get(svc.Name).PkgName
			specs = append(specs, &codegen.ImportSpec{
				Path: path.Join(genpkg, codegen.SnakeCase(svc.Name)),
				Name: pkgName,
			})
			header := codegen.Header("", apiPkg, specs)
			sections = []*codegen.SectionTemplate{header}

			for _, sch := range svc.Schemes {
				data := map[string]interface{}{
					"Scheme":  sch,
					"Service": svc,
				}
				schemeType := strings.ToLower(sch.Type)
				path := fmt.Sprintf("%s_%sauth.go", codegen.SnakeCase(svc.Name), schemeType)
				if _, err := os.Stat(path); !os.IsNotExist(err) {
					break // file already exists, skip it.
				}
				schemeSections := append(sections, &codegen.SectionTemplate{
					Name:   fmt.Sprintf("security-authfunc-%s", schemeType),
					Source: dummyAuthFuncT,
					Data:   data,
				})
				result = append(result, &codegen.File{
					Path:             path,
					SectionTemplates: schemeSections,
					SkipExist:        true,
				})
			}
		}
	}

	return result
}

// input: SchemeData
const dummyAuthFuncT = `
{{ printf "%s%sAuth implements the authorization logic for service %q for the %q security scheme." .Service.StructName .Scheme.Type .Service.Name .Scheme.Type | comment }}
func {{ .Service.StructName }}{{ .Scheme.Type }}Auth(ctx context.Context, {{ if eq .Scheme.Type "Basic" }}user, pass{{ else if eq .Scheme.Type "APIKey" }}key{{ else }}token{{ end }} string, s *security.{{ .Scheme.Type }}Scheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}
`
