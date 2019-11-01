package main

import (
	"context"
	"encoding/json"
	"os"
	"fmt"
	"time"
	"log"
	"net/http"
	"cloud.google.com/go/datastore"
)

const (
	EnvKeyDatastoreProjectId = "DATASTORE_PROJECT_ID"
	EnvKeyPORT               = "PORT"
)

type Entry struct {
	Title     string    `json:"title"'`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv(EnvKeyDatastoreProjectId)
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 最新10件取得
	q := datastore.NewQuery("Entry").Order("-created_at").Limit(10)
	entries := make([]Entry, 0, 10)
	if _, err := client.GetAll(ctx, q, &entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// JSON にして返す
	res, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
