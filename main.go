package main

import (
	//"context"
	//"fmt"
	//"github.com/bellasouzas/travel-cost-api-go/db_config"
	//"github.com/bellasouzas/travel-cost-api-go/routers"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/bellasouzas/travel-cost-api-go/routers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/getalldata", routers.GetAllData)
	http.HandleFunc("/totalcostreservartion", routers.TotalCostReservations)
	http.HandleFunc("/totalcostpermonth", routers.TotalCostPerMonth)
	http.HandleFunc("/findbyname", routers.FindByName)

	log.Fatal(http.ListenAndServe(":7000", nil))

	//
	//
	//conn := db_config.GetMongoConnection()
	//collection := conn.Database("quickstart").Collection("reservations")
	//
	//fmt.Println("Collection accessed!")
	//cursor, err := collection.Find(context.TODO(), bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var travelInfo []bson.M
	//if err = cursor.All(context.TODO(), &travelInfo); err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(travelInfo)

}
