package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thuonghidien/grpc-init/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

// 35.241.78.226:80
// AIzaSyCLuekH90oV-nYyIEmNqK6kYOCyErEPTUc
const (
	address = ":50051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client proto.CustomerClient, customer *proto.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client proto.CustomerClient, filter *proto.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}

		log.Printf("Customer: %v", customer)
	}
}
func main() {

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := proto.NewCustomerClient(conn)

	g := gin.Default()
	g.GET("/create", func(ctx *gin.Context) {

		customer := &proto.CustomerRequest{
			Id:    101,
			Name:  "Shiju Varghese",
			Email: "shiju@xyz.com",
			Phone: "732-757-2923",
			Addresses: []*proto.CustomerRequest_Address{
				{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: false,
				},
				{
					Street:            "Greenfield",
					City:              "Kochi",
					State:             "KL",
					Zip:               "68356",
					IsShippingAddress: true,
				},
			},
		}

		// Create a new customer
		createCustomer(client, customer)
	})

	g.GET("/show", func(ctx *gin.Context) {

		filter := &proto.CustomerFilter{Keyword: ""}
		getCustomers(client, filter)
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
