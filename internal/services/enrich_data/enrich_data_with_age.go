package enrich_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const AgifyURL = "https://api.agify.io/"

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func EnrichDataWithAge(name string) (int, error) {
	url := fmt.Sprintf("%s?name=%s", AgifyURL, name)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
