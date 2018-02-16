package models

import (
	"encoding/json"
	"reflect"
)

type Rate struct {
	Name 				string			`json:"type"`
	StartDate 			string			`json:"start_date"`
	EndDate				string			`json:"end_date"`
	WeekdayRate 		string			`json:"weekday_rate"`
	WeekendRate 		string	 		`json:"weekend_rate"`
}

type RatePayload struct {
	Amount 				string			`json:"Amount"`
	Type 				string			`json:"RateType"`
}

func NewRates() *Rate {
	r := new(Rate)

	return r
}

func (r *Rate) Parse(payload map[string]interface{}) error {

	for k, v := range payload {
		if k == "RateName" {
			r.Name = v.(string)
		} else if k == "StartDate" {
			r.StartDate = v.(string)
		} else if k == "EndDate" {
			r.EndDate = v.(string)
		} else if k == "Rates" {
			for k, v := range v.(map[string]interface{}) {
				if k == "Rate" {
					i := reflect.ValueOf(v.(interface{}))
					switch i.Kind() {
					case reflect.Map:
						rp := new(RatePayload)
						payload := v.(map[string]interface{})
						payloadJson, err := json.Marshal(payload)
						if err != nil {
							return err
						}

						err = json.Unmarshal(payloadJson, &rp)
						if err != nil {
							return err
						}

						if rp.Type == "NIGHTLY_WEEKDAY" {
							r.WeekdayRate = rp.Amount
						} else if rp.Type == "NIGHTLY_WEEKEND" {
							r.WeekendRate = rp.Amount
						}
					case reflect.Slice:
						for _, v := range v.([]interface{}) {
							rp := new(RatePayload)
							payload := v.(map[string]interface{})
							payloadJson, err := json.Marshal(payload)
							if err != nil {
								return err
							}

							err = json.Unmarshal(payloadJson, &rp)
							if err != nil {
								return err
							}

							if rp.Type == "NIGHTLY_WEEKDAY" {
								r.WeekdayRate = rp.Amount
							} else if rp.Type == "NIGHTLY_WEEKEND" {
								r.WeekendRate = rp.Amount
							}
						}
					}
				}
			}
		}
	}

	return nil
}
