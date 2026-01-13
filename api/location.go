package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationItem struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsResponse struct {
	Index []LocationItem `json:"index"`
}

func GetLocations() (LocationsResponse, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"

	resp, err := http.Get(url)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	var data LocationsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationsResponse{}, err
	}

	return data, nil
}
