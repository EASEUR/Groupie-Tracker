package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type geoResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func cleanLocation(loc string) string {
	loc = strings.ReplaceAll(loc, "-", " ")
	loc = strings.ReplaceAll(loc, "_", " ")
	return strings.Title(strings.ToLower(loc))
}

var lastGeocodeCall time.Time

func Geocode(place string) (float64, float64, error) {
	if !lastGeocodeCall.IsZero() {
		elapsed := time.Since(lastGeocodeCall)
		if elapsed < time.Second {
			time.Sleep(time.Second - elapsed)
		}
	}
	lastGeocodeCall = time.Now()

	clean := cleanLocation(place)
	q := url.QueryEscape(clean)

	apiURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		q,
	)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("User-Agent", "Groupie-Tracker")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, 0, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result []geoResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}
	if len(result) == 0 {
		return 0, 0, fmt.Errorf("aucun résultat pour %s", clean)
	}

	lat, _ := strconv.ParseFloat(result[0].Lat, 64)
	lon, _ := strconv.ParseFloat(result[0].Lon, 64)

	return lat, lon, nil
}
