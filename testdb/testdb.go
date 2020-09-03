package testdb

import (
	"context"
	"github.com/bellasouzas/travel-cost-api-go/models"
	"github.com/bellasouzas/travel-mongodb-go/helper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// SeedList creates new values to database to be tested
func SeedReservations() {
	//"DatadeInicio" and "DatadeFim"
	datadeInicio1, _ := time.Parse(time.RFC3339, "1969-12-8T21:00:00-01:00")
	datadeFim1, _ := time.Parse(time.RFC3339, "1969-12-10T21:00:00-05:00")
	datadeInicio2, _ := time.Parse(time.RFC3339, "1969-12-20T21:00:00-02:00")
	datadeFim2, _ := time.Parse(time.RFC3339, "1969-12-23T21:00:00-03:00")

	var reservations = map[string]models.DataTravel{
		"001": {IdentificadorDoProcesso: 15045825, Situacao: "Realizada", ViagemUrgente: "NAO",
			NomeDoOrgaoSuperior: "Ministerio da Educaçao", Nome: "MARINA FERREIRA KITAZONO ANTUNES", Cargo: "",
			DataDeInicio: datadeInicio1, DataDeFim: datadeFim1, Destinos: "Recife/PE", ValorDiarias: 0,
			ValorPassagens: 3406.33, ValorOutrosGastos: 0,
		},
		"002": {
			IdentificadorDoProcesso: 15100682, Situacao: "Realizada", ViagemUrgente: "NAO",
			NomeDoOrgaoSuperior: "Ministerio da Educaçao", Nome: "JORGE ANDRE DE CARVALHO MENDONCA", Cargo: "",
			DataDeInicio: datadeInicio2, DataDeFim: datadeFim2, Destinos: "Recife/PE",
			ValorDiarias: 0, ValorPassagens: 2925.83,
			ValorOutrosGastos: 0},
	}

	collection := helper.ConnectDB()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	for _, datatravel := range reservations {
		if err, _ := collection.InsertOne(ctx, datatravel); err != nil {
			log.Println(err)
		}
	}
}

//Truncate drop database and reinsert the fixture
func Truncate() {
	collection := helper.ConnectDB()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	collection.Drop(ctx)
	SeedReservations()
}

func TestMain(m *testing.M) int {
	//tempDir is a new space to store new datas temporary to mongo
	tempDir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}

	exitVal := m.Run()
	defer os.RemoveAll(tempDir) // clean up
	os.Exit(TestMain(m))
	return exitVal
}
