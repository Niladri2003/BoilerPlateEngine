// routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/niladri2003/BoilerPlateEngine/Controller"
	"github.com/niladri2003/BoilerPlateEngine/engine"
	"github.com/niladri2003/BoilerPlateEngine/models"
	"log"
)

func RegisterRoutes(app *fiber.App, engine *engine.BoilerplateEngine) {
	// Endpoint to generate function signature
	app.Post("/generate", func(c *fiber.Ctx) error {
		var req models.Question
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		log.Printf("Received request: %+v\n", req)
		// Get the list of Questions from the query
		languages := req.Language
		if len(languages) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No languages provided"})
		}
		boilerplates := make(map[string]string)
		for _, lang := range languages {
			resp, err := engine.Generate(models.BoilerplateRequest{Spec: req.ProblemSpec, Language: lang})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			boilerplates[lang] = resp.Signature // Store the generated signature (boilerplate)
		}

		// Create the question object with all the details
		question := models.Question{
			Title:       req.Title,
			Description: req.Description, // You can customize the description
			ProblemSpec: models.ProblemSpec{
				FunctionName: req.ProblemSpec.FunctionName,
				Parameters:   req.ProblemSpec.Parameters,
				ReturnType:   req.ProblemSpec.ReturnType,
				TestCases:    req.ProblemSpec.TestCases,
			},
			Difficulty:   models.Medium, // You can modify this based on actual logic
			Topic:        "General",     // You can also derive this based on the problem
			Language:     languages,     // Save the list of languages
			BoilerPlates: boilerplates,  // Store generated boilerplate code for the requested languages
		}
		log.Println(question)
		if err := Controller.InsetQuestionToDb(question); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		// Send only the signature to frontend
		return c.JSON(fiber.Map{"signature": question.BoilerPlates, "TestCase": question.ProblemSpec.TestCases})
	})

	// Endpoint to handle submission
	app.Post("/submit", func(c *fiber.Ctx) error {
		var submission struct {
			Signature string             `json:"signature"`
			UserCode  string             `json:"user_code"`
			Language  string             `json:"language"`
			Spec      models.ProblemSpec `json:"spec"`
		}
		if err := c.BodyParser(&submission); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Regenerate full boilerplate for merging
		req := models.BoilerplateRequest{Spec: submission.Spec, Language: submission.Language}
		resp, err := engine.Generate(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Merge user code with full boilerplate
		fullCode := engine.MergeSubmission(submission.Signature, submission.UserCode, resp.FullCode)

		// Here, you'd compile/run fullCode against test cases
		return c.SendString("Submission received: " + fullCode)
	})
}
