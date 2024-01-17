package enrich_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const GenderizeURL = "https://api.genderize.io/"

type GenderizeResponse struct {
	Count  int    `json:"count"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func EnrichDataWithGender(name string) (string, error) {
	url := fmt.Sprintf("%s?name=%s", GenderizeURL, name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var genderizeResponse GenderizeResponse
	err = json.Unmarshal(body, &genderizeResponse)
	if err != nil {
		return "", err
	}

	return genderizeResponse.Gender, nil
}
