package api

import (
	"encoding/json"
	"fmt"
	"io"
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

var lastGeocodCall time.Time

func Geocode(place string) (float64, float64, error) {
	if !lastGeocodCall.IsZero() {
		elapsed := time.Since(lastGeocodCall)
		if elapsed < time.Second {
			time.Sleep(time.Second - elapsed)
		}
	}
	lastGeocodCall = time.Now()

	q := url.QueryEscape(place)
	apiURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		q,
	)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("erreur création requête: %v", err)
	}
	// User-Agent requis par Nominatim
	req.Header.Set("User-Agent", "Groupie-Tracker/1.0 (Groupie Tracker App)")

	resp, err := geoClient.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("erreur requête HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Vérifier le statut HTTP
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, 0, fmt.Errorf("API retourned status %d: %s", resp.StatusCode, string(body))
	}

	var result []geoResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, fmt.Errorf("erreur parsing JSON: %v", err)
	}
	if len(result) == 0 {
		return 0, 0, fmt.Errorf("aucun résultat pour: %s", place)
	}

	lat, errLat := strconv.ParseFloat(result[0].Lat, 64)
	lon, errLon := strconv.ParseFloat(result[0].Lon, 64)
	if errLat != nil || errLon != nil {
		return 0, 0, fmt.Errorf("erreur conversion coordonnées: lat=%v, lon=%v", errLat, errLon)
	}
	return lat, lon, nil
}
