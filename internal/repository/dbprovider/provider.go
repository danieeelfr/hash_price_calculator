package dbprovider

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName           = "discountdb"
	UsersCollectionName    = "user"
	ProductsCollectionName = "Product"
)

type DBProvider interface {
	GetUser(id string) error
}

type dbProvider struct {
	DB *mongo.Database
}

func NewDBProvider() (*dbProvider, error) {
	p := new(dbProvider)

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoDb := os.Getenv("MONGO_DB")

	log.Print(mongoUsername)
	log.Print(mongoPassword)
	log.Print(mongoDb)

	// create the mongo context
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect MongoDB
	mongoUri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", mongoUsername, mongoPassword)
	fmt.Println("Connecting to MongoDB...")

	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalf("Error Starting MongoDB Client: %v", err)
	}

	p.DB = client.Database(DatabaseName)

	return p, nil
}

func (ref *dbProvider) GetUser(id string) error {
	c := ref.DB.Collection(UsersCollectionName)
	log.Print(c)

	return nil
}
