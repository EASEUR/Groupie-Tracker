package main

import (
	"fmt"

	"Groupie-Traker/api"
)

func main() {
	// ===== TEST ARTISTS =====
	artists, err := api.GetArtists()
	if err != nil {
		fmt.Println("Erreur artists :", err)
		return
	}
	fmt.Println("Nombre d'artistes :", len(artists))
	if len(artists) > 0 {
		fmt.Println("Premier artiste :", artists[0].Name)
	}

	// ===== TEST LOCATIONS =====
	locations, err := api.GetLocations()
	if err != nil {
		fmt.Println("Erreur locations :", err)
		return
	}
	fmt.Println("Nombre d'entrées locations :", len(locations.Index))

	// ===== TEST DATES =====
	dates, err := api.GetDates()
	if err != nil {
		fmt.Println("Erreur dates :", err)
		return
	}
	fmt.Println("Nombre d'entrées dates :", len(dates.Index))

	// ===== TEST RELATION =====
	relations, err := api.GetRelations()
	if err != nil {
		fmt.Println("Erreur relations :", err)
		return
	}
	fmt.Println("Nombre d'entrées relations :", len(relations.Index))


	// ===== TEST GEOCODE =====
	lat, lon, err := api.Geocode("Paris")
	if err != nil {
    	fmt.Println("Erreur geocode :", err)
	} else {
    	fmt.Printf("Paris : lat=%f, lon=%f\n", lat, lon)
	}

	lat, lon, err = api.Geocode("New_York")
	if err != nil {
    	fmt.Println("Erreur geocode :", err)
	} else {
    	fmt.Printf("New York : lat=%f, lon=%f\n", lat, lon)
	}

	// ===== TEST ARTIST LOCATIONS =====
	if len(relations.Index) > 0 {
    artistID := relations.Index[0].ID
    locs := api.GetArtistLocations(artistID, relations)
    fmt.Printf("Lieux pour l'artiste %d : %v\n", artistID, locs)
    
    // Test avec geocoding des lieux
    for _, loc := range locs {
        lat, lon, err := api.Geocode(loc)
        if err != nil {
            fmt.Printf("- %s : impossible de géocoder\n", loc)
        } else {
            fmt.Printf("- %s : lat=%f, lon=%f\n", loc, lat, lon)
        }
    }
}
}