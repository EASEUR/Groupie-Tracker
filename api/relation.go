package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type RelationItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsResponse struct {
	Index []RelationItem `json:"index"`
}

func GetRelations() (RelationsResponse, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation"

	resp, err := http.Get(url)
	if err != nil {
		return RelationsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RelationsResponse{}, err
	}

	var data RelationsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return RelationsResponse{}, err
	}

	return data, nil
}
