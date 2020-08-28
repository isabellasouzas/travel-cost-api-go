package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bellasouzas/travel-cost-api-go/db_config"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
)

//GetAllData sends a response containing all data on database
func GetAllData(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
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
	fmt.Println("sending response to the client!")
	json.NewEncoder(res).Encode(travelInfo)

}

//TotalCostReservations
func TotalCostReservations(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "TotalCostReservations!\n")
}

//TotalCostPerMonth
func TotalCostPerMonth(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "TotalCostPerMonth!\n")
}

//FindByName
func FindByName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "FindbyName!\n")
}
