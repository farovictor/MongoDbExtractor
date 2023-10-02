package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func transform() {
	// Define the filter to find documents where "age" <= 12
	filter := bson.M{"age": bson.M{"$lte": 12}}

	// Transform the filter into extJSON format
	extJSON, err := bson.MarshalExtJSON(filter, false, false)
	if err != nil {
		fmt.Println("Error converting filter to extJSON:", err)
		return
	}

	// Print the extJSON representation
	fmt.Println(string(extJSON))
}
