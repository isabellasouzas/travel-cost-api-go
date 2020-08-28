package routers

import (
	"encoding/json"
	"fmt"
	"github.com/bellasouzas/travel-cost-api-go/controllers"
	"io"
	"net/http"
)

//GetAllData sends a response containing all data on database
func GetAllData(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Println("Executing GetAllData")
	json.NewEncoder(res).Encode(controllers.GetAllDataFunc())
}

//TotalCostReservations
func TotalCostReservations(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Println("Executing TotalCostReservations")
	json.NewEncoder(res).Encode(controllers.TotalCostReservationsFunc())
}

//TotalCostPerMonth
func TotalCostPerMonth(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "TotalCostPerMonth!\n")
}

//FindByName
func FindByName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "FindbyName!\n")
}
