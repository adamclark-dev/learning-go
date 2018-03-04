package postcode


import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type PostCode struct {
	Status int `json:"status"`
	Result struct {
		Postcode                  string      `json:"postcode"`
		Quality                   int         `json:"quality"`
		Eastings                  int         `json:"eastings"`
		Northings                 int         `json:"northings"`
		Country                   string      `json:"country"`
		NhsHa                     string      `json:"nhs_ha"`
		Longitude                 float64     `json:"longitude"`
		Latitude                  float64     `json:"latitude"`
		EuropeanElectoralRegion   string      `json:"european_electoral_region"`
		PrimaryCareTrust          string      `json:"primary_care_trust"`
		Region                    string      `json:"region"`
		Lsoa                      string      `json:"lsoa"`
		Msoa                      string      `json:"msoa"`
		Incode                    string      `json:"incode"`
		Outcode                   string      `json:"outcode"`
		ParliamentaryConstituency string      `json:"parliamentary_constituency"`
		AdminDistrict             string      `json:"admin_district"`
		Parish                    string      `json:"parish"`
		AdminCounty               interface{} `json:"admin_county"`
		AdminWard                 string      `json:"admin_ward"`
		Ccg                       string      `json:"ccg"`
		Nuts                      string      `json:"nuts"`
		Codes                     struct {
			AdminDistrict             string `json:"admin_district"`
			AdminCounty               string `json:"admin_county"`
			AdminWard                 string `json:"admin_ward"`
			Parish                    string `json:"parish"`
			ParliamentaryConstituency string `json:"parliamentary_constituency"`
			Ccg                       string `json:"ccg"`
			Nuts                      string `json:"nuts"`
		} `json:"codes"`
	} `json:"result"`
}

func Get(postcode string) (PostCode, error) {

	result := PostCode{}
	url := "https://api.postcodes.io/postcodes/" + postcode

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return result, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return result, readErr
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return result, jsonErr
	}

	return result, nil
}
