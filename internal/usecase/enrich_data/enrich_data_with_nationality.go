package enrich_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	// Создаем экземпляр NationalizeResponse для распаковки JSON
	var nationalizeResponse NationalizeResponse

	// Распаковываем JSON в структуру
	err = json.Unmarshal(body, &nationalizeResponse)
	if err != nil {
		fmt.Println("Ошибка при распаковке JSON:", err)
		return "", err
	}
	return nationalizeResponse.Nationality[0].CountryID, nil
}
