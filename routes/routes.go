// routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/niladri2003/BoilerPlateEngine/engine"
	"github.com/niladri2003/BoilerPlateEngine/models"
)

func RegisterRoutes(app *fiber.App, engine *engine.BoilerplateEngine) {
	// Endpoint to generate function signature
	app.Post("/generate", func(c *fiber.Ctx) error {
		var req models.BoilerplateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		resp, err := engine.Generate(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Send only the signature to frontend
		return c.JSON(fiber.Map{"signature": resp.Signature})
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
