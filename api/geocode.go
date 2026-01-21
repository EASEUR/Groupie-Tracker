package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type NominatimResponse []struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func cleanLocation(loc string) string {
	loc = strings.ReplaceAll(loc, "-", " ")
	loc = strings.ReplaceAll(loc, "_", " ")
	return strings.Title(strings.ToLower(loc))
}

func Geocode(place string) (float64, float64, error) {
	clean := cleanLocation(place)

	endpoint := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Add("q", clean)
	params.Add("format", "json")
	params.Add("limit", "1")

	req, err := http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("User-Agent", "groupie-tracker")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var data NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, err
	}

	if len(data) == 0 {
		return 0, 0, fmt.Errorf("aucun résultat pour %s", clean)
	}

	var lat, lon float64
	fmt.Sscanf(data[0].Lat, "%f", &lat)
	fmt.Sscanf(data[0].Lon, "%f", &lon)

	return lat, lon, nil
}
