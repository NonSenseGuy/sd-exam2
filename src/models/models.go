package models

import "time"

type FraudataItem struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	IsReported    bool      `json:"is_reported"`
	ReportReasons []string  `json:"report_reasons,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedOn     time.Time `json:"update_on,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
