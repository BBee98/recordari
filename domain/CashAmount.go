package domain

type CashAmount struct {
	Active      bool        `json:"active"`
	Amount      int64       `json:"amount"`
	Currency    string      `json:"currency"`
	BillingRate BillingRate `json:"billingRate"`
}
