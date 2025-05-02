// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    resource, err := UnmarshalResource(bytes)
//    bytes, err = resource.Marshal()

package domain

import "time"

type Resource []ResourceElement

type ResourceElement struct {
	ID                                     string        `json:"_id"`
	Email                                  Email         `json:"email"`
	Metadata                               string        `json:"metadata"`
	CreatedDate                            time.Time     `json:"createdDate"`
	UpdatedDate                            *time.Time    `json:"updatedDate"`
	Note                                   string        `json:"note"`
	FirstName                              string        `json:"firstName"`
	LastName                               string        `json:"lastName"`
	Status                                 Status        `json:"status"`
	Role                                   Role          `json:"role"`
	IsProjectManager                       bool          `json:"isProjectManager"`
	Links                                  Links         `json:"links"`
	Billing                                Billing       `json:"billing"`
	Billable                               bool          `json:"billable"`
	ResourceRates                          ResourceRates `json:"resourceRates"`
	Tags                                   []interface{} `json:"tags"`
	IsApprover                             bool          `json:"isApprover"`
	CustomFields                           []CustomField `json:"customFields"`
	CalendarIDS                            []CalendarID  `json:"calendarIds"`
	UseCustomAvailability                  bool          `json:"useCustomAvailability"`
	ResourceHasCustomAvailabilityOverrides bool          `json:"resourceHasCustomAvailabilityOverrides"`
	ToilPolicyID                           interface{}   `json:"toilPolicyId"`
}

type Billing struct {
	UseDefault bool  `json:"useDefault"`
	Rate       int64 `json:"rate"`
}

type CustomField struct {
	TemplateID    TemplateID    `json:"templateId"`
	ID            string        `json:"_id"`
	TemplateType  TemplateType  `json:"templateType"`
	TemplateLabel TemplateLabel `json:"templateLabel"`
	Choices       []Choice      `json:"choices"`
}

type Choice struct {
	ID       string `json:"_id"`
	ChoiceID string `json:"choiceId"`
	Value    string `json:"value"`
}

type ResourceRates struct {
	External []interface{} `json:"external"`
	Internal []interface{} `json:"internal"`
}

type CalendarID string

const (
	The680Cc42622001B6Da36115D0 CalendarID = "680cc42622001b6da36115d0"
	The680Cfc77E08Ed345F539Ab1C CalendarID = "680cfc77e08ed345f539ab1c"
)

type TemplateID string

const (
	The680Cfbe8Fdab3F0Db0435068 TemplateID = "680cfbe8fdab3f0db0435068"
	The680Cfbe8Fdab3F0Db0435081 TemplateID = "680cfbe8fdab3f0db0435081"
	The680Cfbe8Fdab3F0Db043508E TemplateID = "680cfbe8fdab3f0db043508e"
)

type TemplateLabel string
type TemplateType string

type Email string

type IconLink string

type Role string

type Status string
