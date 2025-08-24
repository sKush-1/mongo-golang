package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/sKush-1/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		log.Fatal("MONGODB_URL environment variable is not set")
	}

	s, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return s
}
