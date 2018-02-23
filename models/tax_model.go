package models

import (
	"strconv"
)

type TaxFeeDeposit struct {
	Type 			string			`json:"type"`
	Name 			string 			`json:"name"`
	Frequency 		string			`json:"frequency"`
	IsFlat			bool 			`json:"is_flat"`
	Value 			float64 		`json:"value"`
}

func NewTaxFeeDeposit(t string) *TaxFeeDeposit {
	tfd := new(TaxFeeDeposit)
	tfd.Type = t

	return tfd
}

func (t *TaxFeeDeposit) Parse(tax_map map[string]interface{}) error {

	var err error
	for k, v := range tax_map {
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
