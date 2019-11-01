package main

import (
	"context"
	"os"
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

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

type Entries struct {
	Entries []*Entry `json:"entries"`
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv(EnvKeyDatastoreProjectId)
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func indexHandler(gc *gin.Context) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 最新10件取得
	q := datastore.NewQuery("Entry").Order("-created_at").Limit(10)
	entries := make([]*Entry, 0, 10)
	if _, err := client.GetAll(ctx, q, &entries); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, &Entries{Entries: entries})
}

func main() {
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/", indexHandler)

	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)

}
