package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-lodgix-api/clients"
	"github.com/effortless-technologies/elt-lodgix-api/models"

	"github.com/labstack/echo"
)

func GetProperties(c echo.Context) error {

	state := c.QueryParam("state")

	type response struct {
		Count 			int 				`json:"count"`
		Filters 		[]*models.Filter	`json:"filters"`
		Properties		[]*models.Property	`json:"properties"`
	}

	lodgix_resp, err := clients.GetProperties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	lodgix_resp_bytes := []byte(*lodgix_resp)
	properties, err := models.ParseProperties([]byte(lodgix_resp_bytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, filter := range models.Filters {
		if filter.Name == "State" {
			filter.GetOptions(properties)
		}
	}

	resp := new(response)
	//if state == "" {
	//
	//} else {
	//	for _, filter := range models.Filters {
	//		if filter.Name == "State" {
	//			found := false
	//			for _, option := range filter.Options {
	//				if state == option {
	//					found = true
	//					resp.Properties, err = models.FilterState(
	//						properties, state)
	//					if err != nil {
	//						return c.JSON(http.StatusBadRequest, "Bad 'state' query param")
	//					}
	//				}
	//			}
	//
	//			if found == false {
	//				resp.Properties = properties
	//			}
	//		}
	//	}
	//}

	for _, filter := range models.Filters {
		if filter.Name == "State" {
			found := false
			for _, option := range filter.Options {
				if state == option {
					found = true
					resp.Properties, err = models.FilterState(
						properties, state)
					if err != nil {
						return c.JSON(
							http.StatusBadRequest,
							"Bad 'state' query param")
					}
				}
			}

			if found == false {
				resp.Properties = properties
			}
		}
	}


	resp.Count = len(resp.Properties)
	resp.Filters = models.Filters

	return c.JSON(http.StatusAccepted, resp)
}

func GetProps(c echo.Context) error {

	type response struct {
		Count 			int 				`json:"count"`
		Filters 		[]*models.Filter	`json:"filters"`
		Properties		[]*models.Property	`json:"properties"`
	}

	lodgix_resp, err := clients.GetProperties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	lodgix_resp_bytes := []byte(*lodgix_resp)
	properties, err := models.ParseProperties([]byte(lodgix_resp_bytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := new(response)
	//if state == "" {
	//
	//} else {
	//	for _, filter := range models.Filters {
	//		if filter.Name == "State" {
	//			found := false
	//			for _, option := range filter.Options {
	//				if state == option {
	//					found = true
	//					resp.Properties, err = models.FilterState(
	//						properties, state)
	//					if err != nil {
	//						return c.JSON(http.StatusBadRequest, "Bad 'state' query param")
	//					}
	//				}
	//			}
	//
	//			if found == false {
	//				resp.Properties = properties
	//			}
	//		}
	//	}
	//}


	resp.Properties = properties

	resp.Count = len(resp.Properties)
	resp.Filters = models.Filters

	return c.JSON(http.StatusAccepted, resp)
}
