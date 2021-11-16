package models

const MaxListLimit = 200

type IDRequest struct {
	ID string `json:"id"`
}

type ListRequest struct {
	Limit int `json:"limit"`
}

type ReportRequest struct {
	FraudataItem *FraudataItem `json:"fraudata_item"`
}
