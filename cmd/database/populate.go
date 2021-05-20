package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danieeelfr/hash_price_calculator/internal/domain"
	"github.com/danieeelfr/hash_price_calculator/internal/repository/dbprovider"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func populateProductCollection(ctx context.Context, provider *dbprovider.DBProvider) {

	err := provider.DB.Collection("product").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\ncreating products...")
	for i := 1; i < 10; i++ {
		res, err := provider.DB.Collection("product").InsertOne(ctx, domain.Product{
			Id:           primitive.NewObjectID(),
			PriceInCents: i * 5,
			Title:        fmt.Sprintf("Product %v", i),
			Description:  fmt.Sprintf("Description %v", i),
			Discount: &domain.Discount{
				Percentage:   10,
				ValueInCents: i,
			},
		})
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Printf("product created: %s\n", res.InsertedID.(primitive.ObjectID).String())

	}

	fmt.Println("collection discountdb.product populated with success!")
}

func populateUserCollection(ctx context.Context, provider *dbprovider.DBProvider) {

	err := provider.DB.Collection("user").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ncreating users...")
	for i := 1; i < 10; i++ {
		res, err := provider.DB.Collection("user").InsertOne(ctx, domain.User{
			Id:          primitive.NewObjectID(),
			FirstName:   fmt.Sprintf("User %v", i),
			LastName:    fmt.Sprintf("Silva %v", i),
			DateOfBirth: time.Date(1981, 11, i, 0, 0, 0, 0, time.Now().Location()),
		})
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Printf("user created: %s\n", res.InsertedID.(primitive.ObjectID).String())

	}

	fmt.Println("collection discountdb.user populated with success!")
}

func main() {
	ctx := context.Background()
	provider, err := dbprovider.NewDBProvider()
	if err != nil {
		log.Fatal(err)
	}
	populateUserCollection(ctx, provider)
	populateProductCollection(ctx, provider)
}
