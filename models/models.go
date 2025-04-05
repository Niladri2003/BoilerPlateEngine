// models.go
package models

type Parameter struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Subtype string `json:"subtype,omitempty"`
}

type ProblemSpec struct {
	FunctionName string      `json:"function_name"`
	Parameters   []Parameter `json:"parameters"`
	ReturnType   struct {
		Type    string `json:"type"`
		Subtype string `json:"subtype,omitempty"`
	} `json:"return_type"`
	TestCases []map[string]interface{} `json:"test_cases"`
}

type BoilerplateRequest struct {
	Spec     ProblemSpec `json:"spec"`
	Language string      `json:"language"`
}

type BoilerplateResponse struct {
	Signature string `json:"signature"` // Sent to frontend
	FullCode  string `json:"full_code"` // Stored for submission
}
