package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataTravel struct {
	ID                      primitive.ObjectID `bson:"_id,omitemptly"`
	IdentificadorDoProcesso string             `bson:"IdentificadorDoProcesso"`
	Situacao                string             `bson:"Situacao,omitemptly"`
	ViagemUrgente           string             `bson:"ViagemUrgente,omitemptly"`
	NomeDoOrgaoSuperior     string             `bson:"NomeDoOrgao Superior,omitemptly" `
	Nome                    string             `bson:"Nome,omitemptly" `
	Cargo                   string             `bson:"Cargo,omitemptly"`
	DataDeInicio            string             `bson:"DataDeInicio,omitemptly" `
	DataDeFim               string             `bson:"DataDeFim,omitemptly" `
	Destinos                string             `bson:"Destinos,omitemptly"`
	ValorDiarias            float64            `bson:"Valor diarias,omitemptly" `
	ValorPassagens          float64            `bson:"Valor passagens,omitemptly"`
	ValorOutrosGastos       float64            `bson:"Valor outros gastos,omitemptly"`
}