package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//DataTravel modeling database
type DataTravel struct {
	ID                      primitive.ObjectID `bson:"_id,omitemptly"`
	IdentificadorDoProcesso int                `bson:"IdentificadorDoProcesso"`
	Situacao                string             `bson:"Situacao,omitemptly"`
	ViagemUrgente           string             `bson:"ViagemUrgente,omitemptly"`
	NomeDoOrgaoSuperior     string             `bson:"NomeDoOrgao Superior,omitemptly" `
	Nome                    string             `bson:"Nome,omitemptly" `
	Cargo                   string             `bson:"Cargo,omitemptly"`
	DataDeInicio            time.Time          `bson:"Periodo - Data de inicio,omitemptly" `
	DataDeFim               time.Time          `bson:"DataDeFim,omitemptly" `
	Destinos                string             `bson:"Destinos,omitemptly"`
	ValorDiarias            float64            `bson:"Valor diarias,omitemptly" `
	ValorPassagens          float64            `bson:"Valor passagens,omitemptly"`
	ValorOutrosGastos       float64            `bson:"Valor outros gastos,omitemptly"`
	Month                   time.Time          `bson:"month,omitemptly`
}

//ReservationsStatusTotal represents the sum of total cost of reservations
type ReservationsStatusTotal struct {
	ID                    string  `bson:"_id"`
	ReservationsTotalCost float64 `bson:"ReservationsTotalCost"`
}

//TotalCostPerMonth represents total sum of cost of reservations per month
type TotalCostPerMonth struct {
	Month     int     `bson:"month"`
	TotalCost float64 `bson:"totalcost"`
}
