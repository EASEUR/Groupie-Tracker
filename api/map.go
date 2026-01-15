package api

import "fmt"

func PrintArtistMap(artistID int, relations RelationsResponse) {
	locations := GetArtistLocations(artistID, relations)
	if len(locations) == 0 {
		fmt.Println("Pas de lieux trouvés pour cet artiste.")
		return
	}

	fmt.Printf("Lieux de concerts pour l'artiste %d :\n", artistID)
	for _, loc := range locations {
		lat, lon, err := Geocode(loc)
		if err != nil {
			fmt.Printf("- %s : impossible de géocoder\n", loc)
			continue
		}
		fmt.Printf("- %s : lat=%f, lon=%f\n", loc, lat, lon)
	}
}
