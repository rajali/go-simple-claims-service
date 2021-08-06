package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-simple-claims-service/controllers"
	"go-simple-claims-service/driver"
	"go-simple-claims-service/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var claims []models.Claim
var db *sql.DB

func init() {
	gotenv.Load()
	//gotenv.Load(".env.production", "credentials")
}

func main() {
	db = driver.ConnectDB
	controller := controllers.Controller{}

	router := mux.NewRouter()

	//claims = append(claims, models.Claim{ClaimNumber: "C3115610", ClaimSystem: "CC8", PolicySystem: "Exigen", LossDate: "1/1/16", CloseDate: "22/10/19", ClosedOutcome: "Completed", BrandName: "AMI", PolicyType: "Domestic Home", LossCause: "Weather - snow", LossType: "Property", ClaimLineOfBusiness: "Homeowners Line", Coverage: "Base Sum Insured", ClaimDescription: "Damage Spouting (CONSIDERATION ONLY)", GrossIncurred: 2015, NetIncurred: 2015, PcodeCl: 8013},
	//	models.Claim{ClaimNumber: "C3134370", ClaimSystem: "CC8", PolicySystem: "Exigen", LossDate: "1/1/16", CloseDate: "10/09/19", ClosedOutcome: "Completed", BrandName: "AMI", PolicyType: "Domestic Home", LossCause: "Weather - snow", LossType: "Property", ClaimLineOfBusiness: "Homeowners Line", Coverage: "Base Sum Insured", ClaimDescription: "Damage Spouting (CONSIDERATION ONLY)", GrossIncurred: 2015, NetIncurred: 2015, PcodeCl: 8013})

	router.HandleFunc("/claims", controller.GetClaims).Methods("GET")
	router.HandleFunc("/claims/{ClaimNumber}", controller.getClaim).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
