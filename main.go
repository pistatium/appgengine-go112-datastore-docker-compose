package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

const (
	EnvKeyDatastoreProjectId = "DATASTORE_PROJECT_ID"
	EnvKeyPORT               = "PORT"
)

type Entry struct {
	Title     string     `json:"title" binding:"required" datastore:"title,noindex"`
	Body      string     `json:"body" binding:"required" datastore:"body,noindex"`
	CreatedAt *time.Time `json:"created_at" datastore:"created_at"`
}

type Entries struct {
	Entries []*Entry `json:"entries"`
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv(EnvKeyDatastoreProjectId)
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func postEntry(gc *gin.Context) {
	ctx := context.Background()

	var entry Entry
	if err := gc.ShouldBindJSON(&entry); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entry.CreatedAt == nil {
		now := time.Now()
		entry.CreatedAt = &now
	}
	key := datastore.IncompleteKey("Entry", nil)
	if _, err := client.Put(ctx, key, &entry); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
}

func getEntries(gc *gin.Context) {
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

func index(gc *gin.Context) {
	gc.String(http.StatusOK, "try: GET /entries or POST /entries")
}

func main() {
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/", index)
	r.GET("/entries", getEntries)
	r.POST("/entries", postEntry)

	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}
