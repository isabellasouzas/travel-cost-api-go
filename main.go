package main

import (
	"context"
	"fmt"
	"github.com/bellasouzas/travel-cost-api-go/db_config"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	conn := db_config.GetMongoConnection()
	collection := conn.Database("quickstart").Collection("reservations")
	fmt.Println(collection.Find(context.TODO(), bson.M{}))
}
