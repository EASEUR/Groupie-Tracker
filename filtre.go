func FilterArtists(artists []Artist, f Filters) []Artist {
	var result []Artist

	for _, artist := range artists {
		if artist.CreationDate < f.MinCreation || artist.CreationDate > f.MaxCreation {
			continue
		}

		if artist.FirstAlbumDate < f.MinAlbum || artist.FirstAlbumDate > f.MaxAlbum {
			continue
		}

		nbMembers := len(artist.Members)
		if nbMembers < f.MinMembers || nbMembers > f.MaxMembers {
			continue
		}

		if len(f.Locations) > 0 && !HasLocation(artist.Locations, f.Locations) {
			continue
		}

		result = append(result, artist)
	}

	return result
}
