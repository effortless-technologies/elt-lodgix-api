package models

import ()

type Address struct {
	Street1			string		`json:"street_1"`
	Street2 		string		`json:"street_2"`
	City 			string		`json:"city"`
	State 			string 		`json:"state"`
	Country 		string 		`json:"country"`
	ZipCode 		string		`json:"zip_code"`
}

func (a *Address) Parse(addressMap map[string]interface{}) error {

	for k, v := range addressMap {
		if k == "Street1" {
			a.Street1 = v.(string)
		} else if k == "Street2" && v != nil {
			a.Street2 = v.(string)
		} else if k == "City" {
			a.City = v.(string)
		} else if k == "State" {
			for k, v := range v.(map[string]interface{}) {
				if k == "@code" {
					a.State = v.(string)
				}
			}
		} else if k == "Country" {
			for k, v := range v.(map[string]interface{}) {
				if k == "@code" {
					a.Country = v.(string)
				}
			}
		} else if k == "PostalCode" {
			a.ZipCode = v.(string)
		}
	}

	return nil
}
