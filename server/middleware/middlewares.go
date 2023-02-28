package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collector *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectName := os.Getenv("DB_COLLECTION_NAME")
	clienOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clienOption)

	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb")

	client.Database(dbName).Collection(collectName)
}

func Task() {}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {}

func CreateTask() {}

func UndoTask() {}

func DelateTask() {}

func DeleteAllTasks() {}
