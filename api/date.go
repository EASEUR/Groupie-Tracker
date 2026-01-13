package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type DateItem struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesResponse struct {
	Index []DateItem `json:"index"`
}

func GetDates() (DatesResponse, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates"

	resp, err := http.Get(url)
	if err != nil {
		return DatesResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DatesResponse{}, err
	}

	var data DatesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return DatesResponse{}, err
	}

	return data, nil
}
