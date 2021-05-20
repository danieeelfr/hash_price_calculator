package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	DateOfBirth time.Time          `bson:"date_of_birth"`
}

type Product struct {
	Id           primitive.ObjectID `bson:"_id"`
	PriceInCents int                `bson:"price_in_cents"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	Discount     *Discount
}

type Discount struct {
	Percentage   float64 `bson:"percentage"`
	ValueInCents int     `bson:"value_in_cents"`
}
