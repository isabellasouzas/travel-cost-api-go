package routers

import (
	"io"
	"net/http"
)

//GetAllData
func GetAllData(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "GetAllData!\n")
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
