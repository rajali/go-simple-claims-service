package models

type Claim struct {
	ClaimNumber         string `json:"claimNumber"`
	ClaimSystem         string `json:"claimSystem"`
	PolicySystem        string `json:"policySystem"`
	LossDate            string `json:"lossDate"`
	CloseDate           string `json:"closeDate"`
	ClosedOutcome       string `json:"closedOutcome"`
	BrandName           string `json:"brandName"`
	PolicyType          string `json:"policyType"`
	LossCause           string `json:"lossCause"`
	LossType            string `json:"lossType"`
	ClaimLineOfBusiness string `json:"claimLineOfBusiness"`
	Coverage            string `json:"converage"`
	ClaimDescription    string `json:"claimDescription"`
	GrossIncurred       int    `json:"grossIncurred"`
	NetIncurred         int    `json:"netIncurred"`
	PcodeCl             int    `json:"pCodeCl"`
}
