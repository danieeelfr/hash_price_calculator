package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danieeelfr/hash_price_calculator/discountpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := discountpb.NewDiscountServiceClient(cc)

	getDiscount(c)
}

func getDiscount(c discountpb.DiscountServiceClient) {

	input := &discountpb.DiscountRequest{
		UserId:    "123",
		ProductId: "321",
	}

	d, e := c.GetDiscount(context.Background(), input)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(d)

	// res, err := c.CreatePerson(context.Background(), &phonebookpb.PersonRequest{Person: person})
	// if err != nil {
	// 	fmt.Printf("Error while creating the person: %v\n", err)
	// }
	// fmt.Printf("Person Created: %v\n", res)
}
