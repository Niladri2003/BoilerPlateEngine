// engine.go
package engine

import (
	"bytes"
	"fmt"
	"github.com/niladri2003/BoilerPlateEngine/mappings"
	"github.com/niladri2003/BoilerPlateEngine/models"
	"strings"
	"text/template"
)

type BoilerplateEngine struct {
	templates map[string]*template.Template
}

var includes = map[string]string{
	"cpp":    "#include <vector>\nusing namespace std;\n",
	"python": "",
	"java":   "",
}

func NewBoilerplateEngine() (*BoilerplateEngine, error) {
	templates := map[string]*template.Template{
		"cpp":    template.Must(template.ParseFiles("templates/cpp.go.tmpl")),
		"python": template.Must(template.ParseFiles("templates/python.go.tmpl")),
		"java":   template.Must(template.ParseFiles("templates/java.go.tmpl")),
	}
	return &BoilerplateEngine{templates: templates}, nil
}

func (e *BoilerplateEngine) Generate(req models.BoilerplateRequest) (models.BoilerplateResponse, error) {
	lang := req.Language
	spec := req.Spec

	// Map types to language-specific syntax
	params := make([]string, len(spec.Parameters))
	for i, p := range spec.Parameters {
		t := mappings.TypeMappings[lang][p.Type]
		if p.Subtype != "" {
			t = strings.Replace(t, "${subtype}", mappings.TypeMappings[lang][p.Subtype], 1)
		}
		if lang == "cpp" && p.Type == "array" {
			t += "&" // Pass by reference for arrays in C++
		}
		params[i] = fmt.Sprintf("%s %s", t, p.Name)
	}
	returnType := mappings.TypeMappings[lang][spec.ReturnType.Type]
	if spec.ReturnType.Subtype != "" {
		returnType = strings.Replace(returnType, "${subtype}", mappings.TypeMappings[lang][spec.ReturnType.Subtype], 1)
	}

	// Generate signature
	signature := fmt.Sprintf("%s %s(%s)", returnType, spec.FunctionName, strings.Join(params, ", "))
	if lang == "python" {
		signature = fmt.Sprintf("def %s(%s) -> %s", spec.FunctionName, strings.Join(params, ", "), returnType)
	}

	// Generate full boilerplate
	data := struct {
		Includes     string
		ReturnType   string
		FunctionName string
		Parameters   string
		TestCases    []models.TestCase // Now the type is correct
	}{
		Includes:     includes[lang],
		ReturnType:   returnType,
		FunctionName: spec.FunctionName,
		Parameters:   strings.Join(params, ", "),
		TestCases:    spec.TestCases, // This is now compatible
	}

	var fullCode bytes.Buffer
	if err := e.templates[lang].Execute(&fullCode, data); err != nil {
		return models.BoilerplateResponse{}, err
	}

	return models.BoilerplateResponse{
		Signature: signature,
		FullCode:  fullCode.String(),
		TestCase:  spec.TestCases,
	}, nil
}

func (e *BoilerplateEngine) MergeSubmission(signature, userCode, fullCode string) string {
	placeholder := "// Implement here\n"
	if strings.Contains(fullCode, placeholder) {
		return strings.Replace(fullCode, placeholder, userCode, 1)
	}
	return fmt.Sprintf("%s {\n%s\n}", signature, userCode)
}
