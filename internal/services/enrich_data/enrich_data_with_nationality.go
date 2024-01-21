package enrich_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const NationalizeURL = "https://api.nationalize.io/"

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Count       int       `json:"count"`
	Name        string    `json:"name"`
	Nationality []Country `json:"country"`
}

func EnrichDataWithNationality(name string) (string, error) {
	url := fmt.Sprintf(NationalizeURL + "?name=" + name)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var nationalizeResponse NationalizeResponse
	err = json.Unmarshal(body, &nationalizeResponse)
	if err != nil {
		fmt.Println("Ошибка при распаковке JSON:", err)
		return "", err
	}

	return nationalizeResponse.Nationality[0].CountryID, nil
}
