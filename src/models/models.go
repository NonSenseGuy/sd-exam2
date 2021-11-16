package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type FraudataItem struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	IsReported    bool      `json:"is_reported"`
	ReportReasons []string  `json:"report_reasons,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedOn     time.Time `json:"update_on,omitempty"`
}

type FraudataResponseWrapper struct {
	Item  *FraudataItem   `json:"item,omitempty"`
	Items []*FraudataItem `json:"items,omitempty"`
	Code  int             `json:"-"`
}

func (rw *FraudataResponseWrapper) JSON() []byte {
	if rw == nil {
		return []byte("{}")
	}

	res, _ := json.Marshal(rw)

	return res
}

func (rw *FraudataResponseWrapper) StatusCode() int {
	if rw == nil || rw.Code == 0 {
		return http.StatusOK
	}

	return rw.Code
}
