// models.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Difficulty string

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
)

// Question represents a question with parameters and test cases
type Question struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectId, automatically generated
	Title        string             `json:"title"`                   // Title of the question
	Description  string             `json:"description"`             // Description of the question
	ProblemSpec  ProblemSpec        `json:"problem_spec"`            // List of function parameters
	TestCases    []TestCase         `json:"test_cases"`              // List of test cases for this question
	Difficulty   Difficulty         `json:"difficulty"`              // Difficulty level of the question (e.g., Easy, Medium, Hard)
	Topic        string             `json:"topic"`                   // Topic of the question (e.g., Arrays, Trees, DP)
	Language     []string           `json:"language"`                // Programming language for the question (e.g., Python, Java, C++)
	BoilerPlates map[string]string  `json:"boilerplates"`            // Boilerplate code for different languages
}

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
	TestCases []TestCase `json:"test_cases"`
}

type TestCase struct {
	IsPrivate bool                   `json:"is_private"` // Indicates if the test case is private
	Input     map[string]interface{} `json:"input"`      // Key-value pairs for inputs
	Expected  interface{}            `json:"expected"`   // Expected output for this test case
}

type BoilerplateRequest struct {
	Spec     ProblemSpec `json:"spec"`
	Language string      `json:"language"`
}

type BoilerplateResponse struct {
	Signature string     `json:"signature"` // Sent to frontend
	FullCode  string     `json:"full_code"`
	TestCase  []TestCase `json:"test_cases"` // Stored for submission
}
