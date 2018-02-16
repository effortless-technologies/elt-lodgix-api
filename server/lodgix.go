package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-lodgix-api/clients"
	"github.com/effortless-technologies/elt-lodgix-api/models"

	"github.com/labstack/echo"
)

var Filters []*Filter

type Filter struct {
	Name 		string 				`json:"name"`
	Options 	[]string 			`json:"options"`
}

func init() {
	filter := new(Filter)
	filter.Name = "state"
	filter.Options = []string{"All", "CO", "MA"}

	Filters = append(Filters, filter)
}

func GetProperties(c echo.Context) error {

	state := c.QueryParam("state")

	type response struct {
		Count 			int 				`json:"count"`
		Filters 		[]*Filter			`json:"filters"`
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

	if state == "CO" || state == "MA"{
		resp.Properties, err = models.FilterState(properties, state)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Bad 'state' query param")
		}
	} else {
		resp.Properties = properties
	}

	resp.Count = len(resp.Properties)
	resp.Filters = Filters

	return c.JSON(http.StatusAccepted, resp)
}