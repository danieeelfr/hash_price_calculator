package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/danieeelfr/hash_price_calculator/discountpb"
	"google.golang.org/grpc"
)

type GRPCServer interface {
	Start() error
}

type server struct {
}

func NewGRPCServer() (*server, error) {

	return new(server), nil
}

func (ref *server) Start() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	discountpb.RegisterDiscountServiceServer(s, &server{})

	// start a GO Routine
	go func() {
		fmt.Println("Discount Server Started...")
		if err := s.Serve(l); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// wit to exit (Ctrl+C)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block the channel until the signal is received
	<-ch
	fmt.Println("Stopping Discount Server...")
	s.Stop()
	fmt.Println("Closing Listener...")
	l.Close()
	fmt.Println("All done!")
	return nil
}

func (ref *server) GetDiscount(ctx context.Context, req *discountpb.DiscountRequest) (*discountpb.DiscountResponse, error) {

	log.Println(fmt.Sprintf("Received user_id=%s, product_id=%s", req.UserId, req.ProductId))

	return &discountpb.DiscountResponse{
		Discount: 50.50,
	}, nil
	// data := personItem{
	// 	Name:        person.GetName(),
	// 	Email:       person.GetEmail(),
	// 	Phones:      person.GetPhones(),
	// 	LastUpdated: timestamppb.Now(),
	// }
	// res, err := collection.InsertOne(context.Background(), data)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	// }
	// oid, ok := res.InsertedID.(primitive.ObjectID)
	// if !ok {
	// 	return nil, status.Errorf(codes.Internal, "Cannot convert to OID")
	// }
	// data.ID = oid
	// return &phonebookpb.PersonResponse{
	// 	Person: &phonebookpb.Person{
	// 		Id:          data.ID.Hex(),
	// 		Name:        data.Name,
	// 		Email:       data.Email,
	// 		Phones:      data.Phones,
	// 		LastUpdated: data.LastUpdated,
	// 	},
	// }, nil
}
