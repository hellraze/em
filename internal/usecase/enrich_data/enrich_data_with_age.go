package enrich_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const AgifyURL = "https://api.agify.io/"

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func EnrichDataWithAge(name string) (int, error) {
	url := fmt.Sprintf("%s?name=%s", AgifyURL, name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var agifyResponse AgifyResponse
	err = json.Unmarshal(body, &agifyResponse)
	if err != nil {
		return 0, err
	}
	return agifyResponse.Age, nil
}
