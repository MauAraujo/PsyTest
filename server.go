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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type Questionnaire struct {
	Title     string   `json:"testName"`
	Answers   []string `json:"answers"`
	Questions []string `json:"questions"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var credentials Credentials
	var result Credentials

	json.NewDecoder(r.Body).Decode(&credentials)
	fmt.Println(credentials)

	collection := client.Database("psyTest").Collection("users")
	filter := bson.D{{"username", credentials.Username}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	collection.FindOne(ctx, filter).Decode(&result)
	fmt.Println(result.Password)

	if result.Password == "" || credentials.Password != result.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	mySigningKey := []byte(result.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": result.Type == "admin",
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

func CreateTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var questionnaire Questionnaire

	json.NewDecoder(r.Body).Decode(&questionnaire)
	fmt.Println(questionnaire)
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
	router.HandleFunc("/create-questionnaire", CreateTestHandler).Methods(http.MethodPost, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
