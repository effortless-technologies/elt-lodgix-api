package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-lodgix-api/clients"
	"github.com/effortless-technologies/elt-lodgix-api/models"

	"github.com/labstack/echo"
	"fmt"
)

func GetProperties(c echo.Context) error {

	state := c.QueryParam("state")

	type response struct {
		Count 			int 				`json:"count"`
		Filters 		[]*models.Filter	`json:"filters"`
		Properties		[]*models.Property	`json:"properties"`
	}

	lodgixResp, err := clients.GetProperties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	lodgixRespBytes := []byte(lodgixResp)
	properties, err := models.ParseProperties([]byte(lodgixRespBytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, filter := range models.Filters {
		if filter.Name == "State" {
			filter.GetOptions(properties)
		}
	}

	resp := new(response)

	// TODO: handle no state query param tastefully
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

func CreateInquiry(c echo.Context) error {

	propertyId := c.Param("id")
	fmt.Println(propertyId)

	resp, err := clients.PostInquiry(
		propertyId, "Matty", "Berry",
		"matthewberryhill@gmail.com", "US", "559",
		"7559790", "I-would-like-this-property",
		"http://www.effortlessrental.com")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println()
	fmt.Println(resp)
	fmt.Println()

	return c.JSON(http.StatusAccepted, resp)
}
