package models

type Person struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	IsReported    bool     `json:"is_reported"`
	ReportReasons []string `json:report_reasons,omitempty`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
