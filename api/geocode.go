package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type geoResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

var geoClient = http.Client{
	Timeout: 10 * time.Second,
}

func Geocode(place string) (float64, float64, error) {
	q := url.QueryEscape(place)
	apiURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		q,
	)
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", "groupie-tracker")
	resp, err := geoClient.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	var result []geoResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}
	if len(result) == 0 {
		return 0, 0, fmt.Errorf("no result for %s", place)
	}
	lat, _ := strconv.ParseFloat(result[0].Lat, 64)
	lon, _ := strconv.ParseFloat(result[0].Lon, 64)
	return lat, lon, nil
}
