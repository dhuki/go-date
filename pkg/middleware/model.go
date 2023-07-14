package middleware

import "time"

// requestLog - request information from client
type requestLog struct {
	Timestamp     time.Time   `json:"timestamp"`
	CorrelationID interface{} `json:"correlationId"`
	Method        string      `json:"method"`
	URL           string      `json:"url"`
	Status        int         `json:"status"`
	ResponseTime  float64     `json:"responseTime"`
	ResponseSize  int64       `json:"responseSize"`
	ReqBody       interface{} `json:"requestBody"`
}
