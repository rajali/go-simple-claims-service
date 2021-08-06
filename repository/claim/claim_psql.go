package claimRepository

import (
	"github.com/rajali/go-simple-claims-service/models"
	"database/sql"
	"log"
)

type ClaimRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c ClaimRepository) GetClaims(db *sql.DB, claim models.Claim, claims []models.Claim) []models.Claim {
	rows, err := db.Query("select * from claims")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&claim.ClaimNumber, &claim.ClaimSystem, &claim.PolicySystem, &claim.LossDate, &claim.CloseDate, &claim.ClosedOutcome, &claim.BrandName, &claim.PolicyType, &claim.LossCause, &claim.LossType, &claim.ClaimLineOfBusiness, &claim.Coverage, &claim.ClaimDescription, &claim.GrossIncurred, &claim.NetIncurred, &claim.PcodeCl)
		logFatal(err)
		claims = append(claims, claim)
	}

	return claims
}

func (c ClaimRepository) GetClaim(db *sql.DB, claim models.Claim, claimNumber string) models.Claim {
	rows := db.QueryRow("select * from claims where claimNumber=$1", claimNumber)

	err := rows.Scan(&claim.ClaimNumber, &claim.ClaimSystem, &claim.PolicySystem, &claim.LossDate, &claim.CloseDate, &claim.ClosedOutcome, &claim.BrandName, &claim.PolicyType, &claim.LossCause, &claim.LossType, &claim.ClaimLineOfBusiness, &claim.Coverage, &claim.ClaimDescription, &claim.GrossIncurred, &claim.NetIncurred, &claim.PcodeCl)
	logFatal(err)
	return claim
}
