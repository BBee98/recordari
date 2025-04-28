package domain

type BudgetCategory struct {
	BudgetHours      int64  `json:"budgetHours"`
	BudgetCashAmount int64  `json:"budgetCashAmount"`
	BudgetCurrency   string `json:"budgetCurrency"`
	CategoryID       string `json:"categoryId"`
	ID               string `json:"_id"`
}
