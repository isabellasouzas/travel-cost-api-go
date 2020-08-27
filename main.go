package main

import (
	"context"
	"fmt"
	"github.com/bellasouzas/travel-cost-api-go/db_config"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {

	conn := db_config.GetMongoConnection()
	collection := conn.Database("quickstart").Collection("reservations")

	fmt.Println("Collection accessed!")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var travelInfo []bson.M
	if err = cursor.All(context.TODO(), &travelInfo); err != nil {
		log.Fatal(err)
	}

	fmt.Println(travelInfo)

}
