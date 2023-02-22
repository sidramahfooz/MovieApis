package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://sidramahfooz816:sidra123@cluster0.srdik4j.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

// connect with Mongo
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connected")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection is ready")
}

func insertOne(model model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insserted successfully ", inserted.InsertedID)
}

func updateOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
}

func deleteOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", res.DeletedCount)
}

func deleteAll() int64 {
	filter := bson.D{{}}

	res, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", res.DeletedCount)

	return res.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual controller -file

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOne(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkeAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAll()
	json.NewEncoder(w).Encode(count)
}
