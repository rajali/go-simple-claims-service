package controllers

import (
	"rajali.live.com.au/go/go-simple-claims-service/models"
	claimRepository "rajali.live.com.au/go/go-simple-claims-service/repository/claim"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

var claims []models.Claim

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetClaims(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var claim models.Claim
		claims = []models.Claim{}
		claimRepo := claimRepository.ClaimRepository{}
		claims = claimRepo.GetClaims(db, claim, claims)
		json.NewEncoder(w).Encode(claims)
	}
}

func (c Controller) GetClaim(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var claim models.Claim

		claimRepo := claimRepository.ClaimRepository{}
		claim = claimRepo.GetClaim(db, claim, params["ClaimNumber"])

		json.NewEncoder(w).Encode(claim)
	}
}
