BoilerPlateEngine
BoilerPlateEngine is a tool to generate boilerplate code for different programming languages based on a given problem specification. It supports C++, Python, and Java.


Features
Generate function signatures and boilerplate code for C++, Python, and Java.
Supports various data types including integer, string, array, linkedlist, char, boolean, float, and double.
Customizable templates for each language.
Installation
Clone the repository:


git clone https://github.com/niladri2003/BoilerPlateEngine.git
cd BoilerPlateEngine
Install dependencies:


go mod tidy
Usage
Define the problem specification and request in your main.go file:


package main

import (
"fmt"
"github.com/niladri2003/BoilerPlateEngine/engine"
"github.com/niladri2003/BoilerPlateEngine/models"
)

func main() {
engine, err := engine.NewBoilerplateEngine()
if err != nil {
fmt.Println("Error creating engine:", err)
return
}

    req := models.BoilerplateRequest{
        Language: "cpp",
        Spec: models.ProblemSpec{
            FunctionName: "exampleFunction",
            Parameters: []models.Parameter{
                {Name: "param1", Type: "int"},
            },
            ReturnType: models.ReturnType{Type: "void"},
            TestCases:  []map[string]interface{}{},
        },
    }

    resp, err := engine.Generate(req)
    if err != nil {
        fmt.Println("Error generating boilerplate:", err)
        return
    }

    fmt.Println("Generated Signature:", resp.Signature)
    fmt.Println("Generated Code:", resp.FullCode)
}
Run the project:


go run main.go
Project Structure
engine/: Contains the core logic for generating boilerplate code.
mappings/: Contains type mappings for different programming languages.
models/: Contains data models used in the project.
templates/: Contains templates for generating boilerplate code for different languages.
Contributing
Contributions are welcome! Please open an issue or submit a pull request.


sudo docker-compose up -d


License
This project is licensed under the MIT License. See the LICENSE file for details.