package models

const MaxListLimit = 200

type IDRequest struct {
	ID string `json:"id"`
}

type ListRequest struct {
	Limit int `json:"limit"`
}

type ReportRequest struct {
	Item *FraudataItem `json:"fraudata_item"`
}
