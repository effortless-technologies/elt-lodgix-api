package models

import (
	"strconv"
)

type Review struct {
	Title 			string			`json:"title"`
	Date 			string			`json:"date"`
	Reviewer		string			`json:"reviewer"`
	Stars 			int64 			`json:"stars"`
	Description 	string			`json:"description"`
}

func (r *Review) Parse(reviewMap map[string]interface{}) error {

	var err error
	for k, v := range reviewMap {
		if k == "Title" && v != nil {
			r.Title = v.(string)
		} else if k == "Date" {
			r.Date = v.(string)
		} else if k == "Name" {
			r.Reviewer = v.(string)
		} else if k == "Stars" {
			stars := v.(string)
			r.Stars, err = strconv.ParseInt(stars, 10, 32)
			if err != nil {
				return err
			}
		} else if k == "Description" {
			r.Description = v.(string)
		}
	}

	return nil
}
