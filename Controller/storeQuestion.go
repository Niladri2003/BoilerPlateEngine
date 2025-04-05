package Controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/niladri2003/BoilerPlateEngine/db"
	"github.com/niladri2003/BoilerPlateEngine/models"
)

func InsetQuestionToDb(recivedQuestion models.Question) error {
	log.Info(recivedQuestion)
	if db.Client == nil {
		return fmt.Errorf("MongoDB client is not connected")
	}
	log.Info("Inserting question into MongoDB...")
	mongoClient := db.Client
	questionsCollection := mongoClient.Database("problem_solver").Collection("questions")
	// Insert the question document into MongoDB
	_, err := questionsCollection.InsertOne(context.Background(), recivedQuestion)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Sample question inserted successfully!")
	return nil
}
