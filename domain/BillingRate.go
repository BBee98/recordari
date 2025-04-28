package domain

type BillingRate struct {
	UseDefault bool        `json:"useDefault"`
	Rate       int64       `json:"rate"`
	ID         interface{} `json:"id"`
}
