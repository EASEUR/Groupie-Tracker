package api

import (
	"fmt"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FyneWindow() {
	appNew := app.New()
	window := appNew.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(800, 600))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	resultLabel := widget.NewLabel("Les résultats apparaîtront ici")
	resultContainer := container.NewVBox(
		widget.NewCard("Résultats", "", resultLabel),
	)

	resultScroller := container.NewVScroll(resultContainer)

	searchButton := widget.NewButton("Rechercher", func() {
		query := strings.TrimSpace(searchEntry.Text)
		if query == "" {
			resultLabel.SetText("Veuillez entrer un terme de recherche")
			return
		}

		artists, err := GetArtists()
		if err != nil {
			resultLabel.SetText("Erreur récupération artistes: " + err.Error())
			return
		}

		lowerQ := strings.ToLower(query)
		matches := []fyne.CanvasObject{}
		for _, a := range artists {
			if strings.Contains(strings.ToLower(a.Name), lowerQ) {
				info := container.NewVBox(
					widget.NewLabel(fmt.Sprintf("Nom: %s", a.Name)),
					widget.NewLabel(fmt.Sprintf("Membres: %v", a.Members)),
					widget.NewLabel(fmt.Sprintf("Création: %d", a.CreationDate)),
					widget.NewLabel(fmt.Sprintf("Premier album: %s", a.FirstAlbum)),
				)

				locLabel := widget.NewLabel("Cliquer sur 'Voir lieux' pour charger les lieux")
				seeLoc := widget.NewButton("Voir lieux", func(aID int, ll *widget.Label) func() {
					return func() {
						rels, err := GetRelations()
						if err != nil {
							ll.SetText("Erreur récupération relations: " + err.Error())
							return
						}
						locations := GetArtistLocations(aID, rels)
						if len(locations) == 0 {
							ll.SetText("Aucun lieu trouvé pour cet artiste")
							return
						}
						var sb strings.Builder
						for _, loc := range locations {
							lat, lon, err := Geocode(loc)
							if err != nil {
								sb.WriteString(fmt.Sprintf("- %s : géocodage impossible\n", loc))
								continue
							}
							sb.WriteString(fmt.Sprintf("- %s : lat=%.6f lon=%.6f\n", loc, lat, lon))
						}
						ll.SetText(sb.String())
					}
				}(a.ID, locLabel))

				cardContent := container.NewVBox(info, locLabel, seeLoc)
				card := widget.NewCard(a.Name, "", cardContent)
				matches = append(matches, card)
			}
		}

		if len(matches) == 0 {
			resultLabel.SetText("Aucun artiste trouvé pour: " + query)
			resultContainer.Objects = []fyne.CanvasObject{widget.NewCard("Résultats", "", widget.NewLabel("Aucun résultat"))}
			resultContainer.Refresh()
			resultScroller.ScrollToTop()
			return
		}

		resultLabel.SetText(fmt.Sprintf("%d résultat(s) pour: %s", len(matches), query))
		resultContainer.Objects = matches
		resultContainer.Refresh()
		resultScroller.ScrollToTop()
	})

	searchBox := container.NewBorder(
		nil, nil, nil, searchButton,
		searchEntry,
	)

	mainContent := container.NewBorder(
		searchBox, nil, nil, nil,
		resultScroller,
	)

	window.SetContent(mainContent)
	window.ShowAndRun()
}
