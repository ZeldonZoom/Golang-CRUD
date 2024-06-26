package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	model "task1/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb://192.168.2.62:27017"
const dbName = "Company"
const collName = "Employee"

var collection *mongo.Collection

func init() {
	ClientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), ClientOption)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collName)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var employee model.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	inserted, err := collection.InsertOne(context.Background(), employee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One record inserted", inserted)
	json.NewEncoder(w).Encode(inserted)
}

func ReadAllRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	filter := bson.M{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	var AllEmployees []primitive.M

	for cur.Next(context.Background()) {
		var employee bson.M
		err := cur.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		AllEmployees = append(AllEmployees, employee)
	}

	json.NewEncoder(w).Encode(AllEmployees)

}

func UpdateRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)

	params := mux.Vars(r)

	filter := bson.M{"_id": params["id"]}
	update := bson.M{
		"$set": bson.M{
			"gender":    emp.Gender,
			"firstname": emp.FirstName,
			"lastname":  emp.LastName,
			"number":    emp.LastName,
			"salary":    emp.Salary,
			"active":    emp.Active,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params)
	filter := bson.M{"_id": params["id"]}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(deleted)
}
