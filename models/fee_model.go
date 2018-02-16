package models

import (
	"strconv"
)

type Fee struct {
	Name 			string 			`json:"name"`
	Frequency 		string			`json:"frequency"`
	IsFlat			bool 			`json:"is_flat"`
	Type 			string 			`json:"type"`
	Value 			float64 		`json:"value"`
}

func (t *Fee) Parse(fee_map map[string]interface{}) error {

	var err error
	for k, v := range fee_map {
		if k == "Name" {
			t.Name = v.(string)
		} else if k == "Frequency" {
			t.Frequency = v.(string)
		} else if k == "IsFlat" {
			if v.(string) == "No" {
				t.IsFlat = false
			} else if v.(string) == "Yes" {
				t.IsFlat = true
			}
		} else if k == "Type" {
			t.Type = v.(string)
		} else if k == "Value" {
			value := v.(string)
			t.Value, err  = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
		}
	}

	return nil
}