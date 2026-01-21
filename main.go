package main

import (
	"fmt"
	"groupie-tracker/api"
)

func main() {
	//  ===== TEST FYNE =====
	api.FyneWindow()

	
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
	fmt.Println("\n===== GEOCODE =====")
	tests := []string{
		"paris-france",
		"los_angeles-usa",
		"washington_dc-usa",
	}

	for _, city := range tests {
		lat, lon, err := api.Geocode(city)
		if err != nil {
			fmt.Println(city, "→ erreur :", err)
		} else {
			fmt.Printf("%s → lat=%f lon=%f\n", city, lat, lon)
		}
	}

	// ===== TEST ARTIST LOCATIONS =====
	fmt.Println("\n===== ARTIST LOCATIONS =====")
	artistID := 1
	locs := api.GetArtistLocations(artistID, relations)
	fmt.Printf("Lieux de l'artiste %d : %v\n", artistID, locs)

	for _, loc := range locs {
		lat, lon, err := api.Geocode(loc)
		if err != nil {
			fmt.Println(loc, "→ erreur")
		} else {
			fmt.Printf("%s → %f %f\n", loc, lat, lon)
		}
	}

	// ===== TEST FILTRE =====
	fmt.Println("\n===== FILTRE (création >= 2000) =====")
	for _, a := range artists {
		if a.CreationDate >= 2000 {
			fmt.Printf("%s (%d)\n", a.Name, a.CreationDate)
		}
	}

	select {} // garde Fyne ouvert
}

