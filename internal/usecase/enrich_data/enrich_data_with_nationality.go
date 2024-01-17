package enrich_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const NationalizeURL = "https://api.genderize.io/"

type Country struct {
	CountryID   string
	Probability float64
}

type NationalizeResponse struct {
	Count       int       `json:"count"`
	Name        string    `json:"name"`
	Nationality []Country `json:"country"`
}

func EnrichDataWithNationality(name string) (string, error) {
	url := fmt.Sprintf("%s?name=%s", NationalizeURL, name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var nationalizeResponse NationalizeResponse
	err = json.Unmarshal(body, &nationalizeResponse)
	if err != nil {
		return "", err
	}

	return nationalizeResponse.Nationality[0].CountryID, nil
}
