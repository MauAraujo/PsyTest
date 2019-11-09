package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "AQccess-Control-Allow-Headers")

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userResult struct {
		Value float64
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Database("psyTest").Collection("Users").FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&userResult)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(userResult)
	// expectedPassword, ok := users[credentials.Username]

	// if !ok || expectedPassword != credentials.Password {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	sessionToken := uuid.NewV4()

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(120 * time.Second),
	})
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
	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
