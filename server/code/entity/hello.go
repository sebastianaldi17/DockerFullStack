package entity

import "time"

type Pager struct {
	Next  int64 `json:"next"`
	Limit int64 `json:"limit"`
}

type Filter struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Metadata struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	ID        int64     `json:"id"`
}

type HelloLogsRequest struct {
	Pager  Pager  `json:"pager"`
	Filter Filter `json:"filter"`
}

type HelloLogsResponse struct {
	Logs     []Log    `json:"logs"`
	Metadata Metadata `json:"metadata"`
}
