package api

import "strings"

func GetArtistLocations(artistID int, relations RelationsResponse) []string {
	for _, rel := range relations.Index {
		if rel.ID == artistID {
			locations := []string{}
			for loc := range rel.DatesLocations {
				clean := strings.ReplaceAll(loc, "-", " ")
				locations = append(locations, clean)
			}
			return locations
		}
	}
	return nil
}
