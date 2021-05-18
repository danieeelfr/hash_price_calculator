package grpc_server

import (
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

	discountpb.RegisterDiscountServiceServer(s, &ref)

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
