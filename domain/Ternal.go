package domain

type Ternal struct {
	DefaultRateID interface{}   `json:"defaultRateId"`
	CustomRates   []interface{} `json:"customRates"`
}
