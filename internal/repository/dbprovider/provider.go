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
	// DatabaseName           = "discountdb"
	UsersCollectionName    = "user"
	ProductsCollectionName = "Product"
)

type DB interface {
	GetUser(id string) error
}

type DBProvider struct {
	DB *mongo.Database
}

func NewDBProvider() (*DBProvider, error) {
	p := new(DBProvider)

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoDb := os.Getenv("MONGO_DB")

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

	p.DB = client.Database(mongoDb)

	return p, nil
}

func (ref *DBProvider) GetUser(id string) error {
	c := ref.DB.Collection(UsersCollectionName)
	log.Print(c)

	return nil
}
