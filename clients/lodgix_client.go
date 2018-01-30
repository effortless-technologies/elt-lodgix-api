package clients

import (
	"net/http"

	"io/ioutil"
)

func GetProperties() (*string, error) {

	url := "http://www.lodgix.com/api/xml/properties/get?" +
		"Token=5b3c8006942b12b1f214a1bbd7bcb5e5" +
			"&IncludePhotos=Yes" +
				"&OwnerID=27102" +
					"&IncludeRates=Yes" +
						"&JSONCallback=1" +
							"&IncludeAmenities=Yes" +
								"&IncludeTaxes=Yes" +
									"&IncludeReviews=Yes" +
										"&IncludeMergedRates=Yes"

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)

	return &bodyString, nil
}

// http://www.lodgix.com/api/xml/properties/get?Token=5b3c8006942b12b1f214a1bbd7bcb5e5&IncludePhotos=Yes&OwnerID=27102&IncludeRates=Yes&JSONCallback=1&IncludeAmenities=Yes&IncludeTaxes=Yes&IncludeReviews=Yes
