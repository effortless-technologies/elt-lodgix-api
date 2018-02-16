package models

import (
	"encoding/json"
	"reflect"
	"strconv"
	"fmt"
)

type Property struct {
	Name 				string				`json:"name"`
	LogdixId			int					`json:"logdix_id"`
	Address 			*Address			`json:"address"`
	Bedrooms			int					`json:"bedrooms"`
	Bathrooms			float64				`json:"bathrooms"`
	Sleeps				int					`json:"sleeps"`
	Rates 				[]*Rate				`json:"rates"`
	Taxes 				[]*TaxFeeDeposit 	`json:"taxes"`
	Fees 				[]*TaxFeeDeposit	`json:"fees"`
	Deposits 			[]*TaxFeeDeposit 	`json:"deposits"`
	CheckIn				string				`json:"check_in"`
	CheckOut			string 				`json:"check_out"`
	Description 		string				`json:"description"`
	MarketingTitle 		string				`json:"marketing_title"`
	MarketingTeaser		string				`json:"marketing_teaser"`
	Amenities 			[]string			`json:"amenities"`
	Type				string				`json:"type"`
	Images 				[]*Image			`json:"images"`
	Reviews				[]*Review			`json:"reviews"`
}

func NewProperty() *Property {
	p := new(Property)

	return p
}

func (p *Property) Parse(property_map map[string]interface{}) error {

	var err error
	for k, v := range property_map {
		if k == "Name" {
			p.Name = v.(string)
		} else if k == "ID" {
			lodgixId := v.(string)
			p.LogdixId, err = strconv.Atoi(lodgixId)
			if err != nil {
				return err
			}
		} else if k == "Address" {
			address := new(Address)
			address.Parse(v.(map[string]interface{}))
			p.Address = address
		} else if k == "Bedrooms" {
			bedrooms := v.(string)
			p.Bedrooms, err = strconv.Atoi(bedrooms)
			if err != nil {
				return err
			}
		} else if k == "Baths" {
			bathrooms := v.(string)
			p.Bathrooms, err = strconv.ParseFloat(bathrooms, 64)
			if err != nil {
				return err
			}
		} else if k == "MaxGuests" {
			sleeps := v.(string)
			p.Sleeps, err = strconv.Atoi(sleeps)
			if err != nil {
				return err
			}
		} else if k == "CheckIn" {
			p.CheckIn = v.(string)
		} else if k == "CheckOut"{
			p.CheckOut = v.(string)
		} else if k == "MarketingTitle" {
			p.MarketingTitle = v.(string)
		} else if k == "MarketingTeaser" {
			p.MarketingTeaser = v.(string)
		} else if k == "Description" {
			p.Description = v.(string)
		} else if k == "MergedRates" {
			for k, v := range v.(map[string]interface{}) {
				if k == "RatePeriod" {
					i := reflect.ValueOf(v.(interface{}))
					switch i.Kind() {
					case reflect.Map:
						rate := new(Rate)
						rate.Parse(v.(map[string]interface{}))
						p.Rates = append(p.Rates, rate)
					case reflect.Slice:
						for _, v := range v.([]interface{}) {
							rate := new(Rate)
							rate.Parse(v.(map[string]interface{}))
							p.Rates = append(p.Rates, rate)
						}
					}
				}
			}
		} else if k == "TaxesFeesDeposits" {
			for k, v := range v.(map[string]interface{}) {
				if k == "Taxes" && v != nil {
					for k, v := range v.(map[string]interface{}) {
						if k == "Tax" && v != nil {
							i := reflect.ValueOf(v.(interface{}))
							switch i.Kind() {
							case reflect.Map:
								tax := NewTaxFeeDeposit("tax")
								tax.Parse(v.(map[string]interface{}))
								p.Taxes = append(p.Taxes, tax)
							case reflect.Slice:
								for _, v := range v.([]interface{}) {
									tax := NewTaxFeeDeposit("tax")
									tax.Parse(v.(map[string]interface{}))
									p.Taxes = append(p.Taxes, tax)
								}
							}
						}
					}
				} else if k == "Fees" && v != nil {
					for k, v := range v.(map[string]interface{}) {
						if k == "Fee" && v != nil {
							i := reflect.ValueOf(v.(interface{}))
							switch i.Kind() {
							case reflect.Map:
								fee := NewTaxFeeDeposit("fee")
								fee.Parse(v.(map[string]interface{}))
								p.Fees = append(p.Fees, fee)
							case reflect.Slice:
								for _, v := range v.([]interface{}) {
									fee := NewTaxFeeDeposit("fee")
									fee.Parse(v.(map[string]interface{}))
									p.Fees = append(p.Fees, fee)
								}
							}
						}
					}
				} else if k == "Deposits" && v != nil {
					for k, v := range v.(map[string]interface{}) {
						if k == "Deposit" && v != nil {
							i := reflect.ValueOf(v.(interface{}))
							switch i.Kind() {
							case reflect.Map:
								deposit := NewTaxFeeDeposit("deposit")
								deposit.Parse(v.(map[string]interface{}))
								p.Deposits = append(p.Deposits, deposit)
							case reflect.Slice:
								for _, v := range v.([]interface{}) {
									deposit := NewTaxFeeDeposit("deposit")
									deposit.Parse(v.(map[string]interface{}))
									p.Deposits = append(p.Deposits, deposit)
								}
							}
						}
					}
				}
			}
		} else if k == "Amenities" {
			for k, v := range v.(map[string]interface{}) {
				if k == "Amenity" {
					for _, v := range v.([]interface{}) {
						amenity := v.(map[string]interface{})
						for k, v := range amenity {
							if k == "Name" {
								p.Amenities = append(p.Amenities, v.(string))
							}
						}
					}
				}
			}
		} else if k == "PropertyType" {
			p.Type = v.(string)
		} else if k == "Photos" {
			for k, v := range v.(map[string]interface{}) {
				if k == "Photo" {
					for _, v := range v.([]interface{}) {
						image := new(Image)
						image.Parse(v.(map[string]interface{}))
						p.Images = append(p.Images, image)
					}
				}
			}
		} else if k == "Reviews" {
			for k, v := range v.(map[string]interface{}) {
				if k == "Review" {
					for _, v := range v.([]interface{}) {
						review := new(Review)
						review.Parse(v.(map[string]interface{}))
						p.Reviews = append(p.Reviews, review)
					}
				}
			}
		}
	}

	return nil
}

func ParseProperties(payload []byte) ([]*Property, error) {

	i := 0
	payload = append(payload[:i], payload[i+1:]...)
	payload = append(payload[:i], payload[i+1:]...)
	payload = payload[:len(payload)-2]

	results := make(map[string]interface{})
	err := json.Unmarshal(payload, &results)
	if err != nil {
		panic(err)
	}

	var results_array []interface{}
	for _, v := range results {
		for k, v := range v.(map[string]interface{}) {
			if k == "Properties" {
				for k, v := range v.(map[string]interface{}) {
					if k == "Property" {
						results_array = v.([]interface{})
					}
				}
			}
		}
	}

	properties := []*Property{}
	for _, v := range results_array {
		property_map := v.(map[string]interface{})
		property := NewProperty()
		err := property.Parse(property_map)
		if err != nil {
			return nil, err
		}

		properties = append(properties, property)
	}

	return properties, nil
}

func FilterState(properties []*Property, state string) ([]*Property, error) {

	var filtered []*Property
	for _, p := range properties {
		if p.Address.State == state {
			filtered = append(filtered, p)
		}
	}

	fmt.Println()
	fmt.Println(len(filtered))
	fmt.Println()

	return filtered, nil
}
