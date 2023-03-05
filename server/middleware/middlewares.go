package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	TastComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

func TaskComplete(task string) {

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result)

}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Control-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Origin", "*")
	w.Header().Set("Access-Control-Methods", "PUT")
	w.Header().Set("Access-Control-Headers", "Content-Type")

	params := mux.Vars(r)

	UndoTask(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Control-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Origin", "*")

	td := deleteAlltask()

	json.NewEncoder(w).Encode(td)
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}
func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	curl.Close()

	return results
}

func deleteOneTask(task Todo) {}

func deleteAllTask(task Todo) {}
