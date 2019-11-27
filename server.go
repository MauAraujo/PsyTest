package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	FirstName  string             `json:"firstName"`
	MiddleName string             `json:"middleName"`
	LastName   string             `json:"lastName"`
	Age        int                `json:"age"`
	Gender     string             `json:"gender"`
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Type       string             `json:"type"`
	ID         primitive.ObjectID `bson:"_id,omitempty"`
}

type Questionnaire struct {
	Title       string             `json:"testName"`
	Description string             `json:"description"`
	Answers     []string           `json:"answers"`
	Questions   []string           `json:"questions"`
	Uid         primitive.ObjectID `bson:"_uid,omitempty"`
	ID          primitive.ObjectID `bson:"_id,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User
	var result User
	json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("psyTest").Collection("Users")
	filter := bson.D{{"username", user.Username}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.FindOne(ctx, filter).Decode(&result)

	fmt.Println(result)

	if result.Password == "" || user.Password != result.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	mySigningKey := []byte(result.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": result.Type == "admin",
		"uid":   result.ID,
	})
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: time.Now().Add(120 * time.Second),
	})
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	collection := client.Database("psyTest").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.InsertOne(ctx, bson.M{
		"firstName":  user.FirstName,
		"middleName": user.MiddleName,
		"lastName":   user.LastName,
		"age":        user.Age,
		"gender":     user.Gender,
		"type":       user.Type,
		"username":   user.Username,
		"password":   user.Password})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	collection := client.Database("psyTest").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var results []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	resJson, _ := json.Marshal(results)
	w.Write(resJson)
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	collection := client.Database("psyTest").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.InsertOne(ctx, bson.M{
		"firstName":  user.FirstName,
		"middleName": user.MiddleName,
		"lastName":   user.LastName,
		"age":        user.Age,
		"gender":     user.Gender,
		"type":       user.Type,
		"username":   user.Username,
		"password":   user.Password})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	collection := client.Database("psyTest").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.InsertOne(ctx, bson.M{
		"firstName":  user.FirstName,
		"middleName": user.MiddleName,
		"lastName":   user.LastName,
		"age":        user.Age,
		"gender":     user.Gender,
		"type":       user.Type,
		"username":   user.Username,
		"password":   user.Password})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func CreateQuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var questionnaire Questionnaire

	json.NewDecoder(r.Body).Decode(&questionnaire)
	fmt.Println(questionnaire)

	collection := client.Database("psyTest").Collection("Questionnaires")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	uid, _ := primitive.ObjectIDFromHex(questionnaire.Uid.Hex())

	_, err := collection.InsertOne(ctx, bson.M{"_uid": uid, "testName": questionnaire.Title, "description": questionnaire.Description, "questions": questionnaire.Questions, "answers": questionnaire.Answers})
	if err != nil {
		log.Println("Error creating new questionnaire")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetQuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	collection := client.Database("psyTest").Collection("Questionnaires")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Error querying questionnaires")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var results []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	resJson, _ := json.Marshal(results)
	w.Write(resJson)
}

func DeleteQuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := mux.Vars(r)

	qid, _ := primitive.ObjectIDFromHex(v["qid"])
	fmt.Println("Deleting " + qid.String())
	collection := client.Database("psyTest").Collection("Questionnaires")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", qid}}

	deleteResult, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Deleted %v documents in the questionnaires collection\n", deleteResult.DeletedCount)
}

func EditQuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := mux.Vars(r)
	qid, _ := primitive.ObjectIDFromHex(v["qid"])

	var questionnaire Questionnaire

	collection := client.Database("psyTest").Collection("Questionnaires")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", qid}}

	json.NewDecoder(r.Body).Decode(&questionnaire)
	uid, _ := primitive.ObjectIDFromHex(questionnaire.Uid.Hex())

	update := bson.D{
		{"$set", bson.M{"_uid": uid, "testName": questionnaire.Title, "description": questionnaire.Description, "questions": questionnaire.Questions, "answers": questionnaire.Answers}},
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating new questionnaire")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func GetUserQuestionnairesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := mux.Vars(r)
	uid, _ := primitive.ObjectIDFromHex(v["uid"])

	collection := client.Database("psyTest").Collection("Questionnaires")
	filter := bson.D{{"_uid", uid}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("Error querying user questionnaires")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var results []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	// var results []Questionnaire
	// cur.All(ctx, &results)

	resJson, _ := json.Marshal(results)
	w.Write(resJson)

}

func UserQuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := mux.Vars(r)
	qid := v["qid"]

	collection := client.Database("psyTest").Collection("Questionnaires")
	id, _ := primitive.ObjectIDFromHex(qid)
	filter := bson.D{{"_id", id}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if r.Method == "GET" {
		dbresult := collection.FindOne(ctx, filter)
		var result bson.M
		err := dbresult.Decode(&result)
		questionnaire, err := json.Marshal(result)

		if err != nil {
			log.Println("Error getting questionnaire")
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(questionnaire)
	} else if r.Method == "DELETE" {
		deleteResult, err := collection.DeleteOne(ctx, filter)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Deleted %v documents in the questionnaires collection\n", deleteResult.DeletedCount)
	}
}

func initDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mclient, err := mongo.Connect(ctx, options.Client().
		ApplyURI("mongodb://192.168.100.44:27017,192.168.100.45:27018,192.168.100.46:27019/?replicaSet=rs0&readPreference=nearest"))

	if err != nil {
		log.Println("Error connecting to MongoDB")
		return
	}
	client = mclient
}

func main() {
	initDB()
	router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users/get", GetUserHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/create", CreateUserHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users/edit/{uid}", EditUserHandler).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/users/delete/{uid}", DeleteUserHandler).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/questionnaires/get", GetQuestionnaireHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/questionnaires/create", CreateQuestionnaireHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/questionnaires/edit/{qid}/", EditQuestionnaireHandler).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/questionnaires/delete/{qid}/", DeleteQuestionnaireHandler).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/users/{uid}/questionnaires", GetUserQuestionnairesHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/{uid}/questionnaires/{qid}", UserQuestionnaireHandler).Methods(http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPost, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))
	fmt.Println("Server listening on port 3000...")
	log.Println(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
