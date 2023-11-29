package main

import (
	"example-engine-go/pkg/orderprocessor"
	order "example-engine-go/proto"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	// Example order
	orderProto := &order.Order{
		Company: "ABC Company",
		Payment: &order.Payment{
			Type:   "credit",
			Amount: 100.50,
		},
		Products: []*order.Product{
			{
				Name:          "orange",
				Quantity:      3,
				UnitPrice:     30,
				UnitDiscount:  0,
				TotalPrice:    90,
				TotalDiscount: 0,
				Sku:           123,
			},
			{
				Name:          "bag",
				Quantity:      1,
				UnitPrice:     10.50,
				UnitDiscount:  0,
				TotalPrice:    10.50,
				TotalDiscount: 0,
				Sku:           124,
			},
		},
	}

	// Convert Order to []byte
	orderData, err := proto.Marshal(orderProto)
	if err != nil {
		fmt.Println("Error marshalling Protocol Buffers data:", err)
		return
	}

	// Insert the order and get the order number
	orderNumber, err := orderprocessor.InsertOrder(orderData)
	if err != nil {
		fmt.Println("Error processing order:", err)
		return
	}

	fmt.Printf("Order successfully processed with order number: %s\n", orderNumber)
}
