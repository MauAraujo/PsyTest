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
	Username string             `json:"username"`
	Password string             `json:"password"`
	Type     string             `json:"type"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
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
		log.Fatal("Error creating new questionnaire")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
		log.Fatal("Error querying user questionnaires")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var results []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	// var results []Questionnaire
	// cur.All(ctx, &results)

	resJson, _ := json.Marshal(results)
	w.Write(resJson)

}

func QuestionnaireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := mux.Vars(r)
	qid := v["qid"]

	var result Questionnaire
	collection := client.Database("psyTest").Collection("Questionnaires")
	id, _ := primitive.ObjectIDFromHex(qid)
	filter := bson.D{{"_id", id}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if r.Method == "GET" {
		collection.FindOne(ctx, filter).Decode(&result)
		fmt.Println(result)
	} else if r.Method == "DELETE" {
		deleteResult, err := collection.DeleteOne(ctx, filter)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Deleted %v documents in the questionnaires collection\n", deleteResult.DeletedCount)
	}
}

func initDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mclient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal("Error connecting to MongoDB")
		return
	}
	client = mclient
}

func main() {
	initDB()
	router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/create-questionnaire", CreateQuestionnaireHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/user/{uid}/questionnaires", GetUserQuestionnairesHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/user/{uid}/questionnaires/{qid}", QuestionnaireHandler).Methods(http.MethodGet, http.MethodDelete, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))
	fmt.Println("Server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
