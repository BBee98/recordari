package domain

type Budget struct {
	HasBudget    bool       `json:"hasBudget"`
	ProjectHours Hours      `json:"projectHours"`
	CashAmount   CashAmount `json:"cashAmount"`
}
