package enrich_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const GenderizeURL = "https://api.genderize.io/"

type GenderizeResponse struct {
	Count  int    `json:"count"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func EnrichDataWithGender(name string) (string, error) {
	url := fmt.Sprintf("%s?name=%s", GenderizeURL, name)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
