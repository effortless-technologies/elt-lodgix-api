package clients

import (
	"io/ioutil"
	"net/http"
)

func GetProperties() (string, error) {

	url := "http://www.lodgix.com/api/xml/properties/get?" +
		"Token=9e1676ac6c97efef555cdbe817f28a92" +
			"&IncludePhotos=Yes" +
				"&OwnerID=27616" +
					"&IncludeRates=Yes" +
						"&JSONCallback=1" +
							"&IncludeAmenities=Yes" +
								"&IncludeTaxes=Yes" +
									"&IncludeReviews=Yes" +
										"&IncludeMergedRates=Yes"

	var client http.Client
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "Failed making request", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Failed reading response body", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}

// http://www.lodgix.com/api/xml/properties/get?Token=5b3c8006942b12b1f214a1bbd7bcb5e5&IncludePhotos=Yes&OwnerID=27102&IncludeRates=Yes&JSONCallback=1&IncludeAmenities=Yes&IncludeTaxes=Yes&IncludeReviews=Yes

func PostInquiry(
	id string, first_name string, last_name string, email string,
	countryCode string, areaCode string, phoneNumber string,
	message string, referringUrl string)(string, error) {

	url := "http://www.lodgix.com/api/xml/properties/inquirie?" +
		"Token=9e1676ac6c97efef555cdbe817f28a92" +
			"&PropertyID=" + id +
				"&FirstName=" + first_name +
					"&LastName=" + last_name +
						"&Email=" + email +
							"&CountryCode=" + countryCode +
 								"&AreaCode=" + areaCode +
									"&PhoneNumber=" + phoneNumber +
										"&Message=" + message +
											"&ReferringURL=" + referringUrl

	var client http.Client
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "Failed making request", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Failed reading response body", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
