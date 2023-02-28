package middleware

import (
	"context"
	"encoding/json"
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

	collection := client.Database(dbName).Collection(collectName)
	fmt.Println("Collection instance created")
}

func Task() {}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)

}

func CreateTask( w http.ResponseWriter, r *http.Request)) {}

func UndoTask() {}

func DelateTask() {}

func DeleteAllTasks() {}
