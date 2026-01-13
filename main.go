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
}
