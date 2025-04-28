package domain

import (
	"time"
)

type Project struct {
	ID                       string           `json:"_id"`
	Name                     string           `json:"name"`
	Links                    Links            `json:"links"`
	Note                     string           `json:"note"`
	CreatedDate              time.Time        `json:"createdDate"`
	UpdatedDate              time.Time        `json:"updatedDate"`
	TimeEntryEnabled         bool             `json:"timeEntryEnabled"`
	TimeEntryLocked          bool             `json:"timeEntryLocked"`
	TimeEntryApproval        bool             `json:"timeEntryApproval"`
	ProjectCode              string           `json:"projectCode"`
	Tags                     []interface{}    `json:"tags"`
	TimeEntryNoteRequired    bool             `json:"timeEntryNoteRequired"`
	WorkDays                 []bool           `json:"workDays"`
	UseProjectDays           bool             `json:"useProjectDays"`
	BudgetCategories         []BudgetCategory `json:"budgetCategories"`
	FixedCosts               []interface{}    `json:"fixedCosts"`
	Budget                   Budget           `json:"budget"`
	BudgetHours              int64            `json:"budgetHours"`
	BudgetCashAmount         int64            `json:"budgetCashAmount"`
	BudgetCurrency           string           `json:"budgetCurrency"`
	CompanyBillingRateID     interface{}      `json:"companyBillingRateId"`
	ResourceRates            []interface{}    `json:"resourceRates"`
	UseStatusColor           bool             `json:"useStatusColor"`
	Status                   string           `json:"status"`
	UseProjectDuration       bool             `json:"useProjectDuration"`
	Start                    interface{}      `json:"start"`
	End                      interface{}      `json:"end"`
	IncludeBookedTimeReports bool             `json:"includeBookedTimeReports"`
	IncludeBookedTimeGrid    bool             `json:"includeBookedTimeGrid"`
	Private                  bool             `json:"private"`
	ProjectManagers          []interface{}    `json:"projectManagers"`
	Resources                []interface{}    `json:"resources"`
	BackgroundColor          string           `json:"backgroundColor"`
	Metadata                 string           `json:"metadata"`
	CustomFields             []interface{}    `json:"customFields"`
	Billable                 bool             `json:"billable"`
	ProjectRate              Rate             `json:"projectRate"`
	Customers                []interface{}    `json:"customers"`
	CategoryGroups           []string         `json:"categoryGroups"`
	DefaultCategory          string           `json:"defaultCategory"`
}
