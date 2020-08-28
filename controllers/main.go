package controllers

import (
	"context"
	"fmt"
	"github.com/bellasouzas/travel-cost-api-go/db_config"
	"github.com/bellasouzas/travel-cost-api-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

//GetAllData sends a response containing all data on database
func GetAllDataFunc() []bson.M {
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
	return travelInfo
}

//TotalCostReservationsFunc sends a response containing the sum of total reservations cost
func TotalCostReservationsFunc() []models.ReservationsStatusTotal {
	conn := db_config.GetMongoConnection()
	collection := conn.Database("quickstart").Collection("reservations")
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	pipelineResult, pipeline, groupStage := make([]models.ReservationsStatusTotal, 0), make([]bson.M, 0), bson.M{
		"$group": bson.M{
			"_id":                   "$ID",
			"ReservationsTotalCost": bson.M{"$sum": "$Valor diarias"}},
	}

	pipeline = append(pipeline, groupStage)

	data, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
		return nil
	}

	err = data.All(ctx, &pipelineResult)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to decode results", err.Error())
		return nil
	}

	return pipelineResult
}

//TotalCostPerMonthFunc
func TotalCostPerMonthFunc(month int)  ([]models.TotalCostPerMonth) {
//TODO validar input recebido, o parametro precisa ser um numero entre 1 a 12
	conn := db_config.GetMongoConnection()
	collection := conn.Database("quickstart").Collection("reservations")
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	pipelineResult := make([]models.TotalCostPerMonth, 0)
	pipeline := make([]bson.M, 0)
	groupStage:= bson.M{
		"$group": bson.M{
			"month": "$month",
			"TotalCost": bson.M{"$sum": "$expr", "$month": bson.M{
				"$eq": "$Periodo - Data de inicio"}, 1}},
		},
	}

	pipeline = append(pipeline, groupStage)

	data, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
		return nil
	}

	err = data.All(ctx, &pipelineResult)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to decode results", err.Error())
		return nil
	}

	return pipelineResult

}