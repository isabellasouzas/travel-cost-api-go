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

//TotalCostReservations sends a response containing the total sum of Reservations
func TotalCostReservations(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Println("Executing TotalCostReservations")
	json.NewEncoder(res).Encode(controllers.TotalCostReservationsFunc())
}

//FindByName sends a response containing the total cost reservation by name
func FindByName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "FindbyName!\n")
}
