package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FyneWindow() {
	appNew := app.New()
	window := appNew.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(1000, 700))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	// Filtres dynamiques
	memberCountEntry := widget.NewEntry()
	memberCountEntry.SetPlaceHolder("Min nombre de membres...")

	creationYearEntry := widget.NewEntry()
	creationYearEntry.SetPlaceHolder("Année de création (ex: 1990)...")

	resultContainer := container.NewVBox()
	resultScroller := container.NewVScroll(resultContainer)

	var searchTimer *time.Timer
	// Fonction de recherche dynamique avec filtres
	performSearch := func(query string, minMembers string, creationYear string) {
		query = strings.TrimSpace(query)

		artists, err := GetArtists()
		if err != nil {
			log.Printf("❌ Erreur API GetArtists: %v\n", err)
			resultContainer.Objects = []fyne.CanvasObject{
				widget.NewCard("Erreur", "", widget.NewLabel("Erreur: "+err.Error())),
			}
			resultContainer.Refresh()
			return
		}

		lowerQ := strings.ToLower(query)
		matches := []Artist{}

		// Filtrer les artistes selon les critères
		for _, a := range artists {
			// Filtrer par nom
			if query != "" && !strings.Contains(strings.ToLower(a.Name), lowerQ) {
				continue
			}

			// Filtrer par nombre de membres
			if minMembers != "" {
				min, err := strconv.Atoi(strings.TrimSpace(minMembers))
				if err == nil && len(a.Members) < min {
					continue
				}
			}

			// Filtrer par année de création
			if creationYear != "" {
				year, err := strconv.Atoi(strings.TrimSpace(creationYear))
				if err == nil && a.CreationDate != year {
					continue
				}
			}

			matches = append(matches, a)
		}

		// Afficher les résultats
		if len(matches) == 0 {
			message := "Aucun résultat"
			if query != "" || minMembers != "" || creationYear != "" {
				message = "Aucun artiste ne correspond aux critères"
			}
			resultContainer.Objects = []fyne.CanvasObject{
				widget.NewCard("Résultats", "", widget.NewLabel(message)),
			}
			resultContainer.Refresh()
			resultScroller.ScrollToTop()
			return
		}

		// Construire l'affichage des résultats
		cards := make([]fyne.CanvasObject, 0, len(matches))

		// Titre avec nombre de résultats
		titleText := fmt.Sprintf("Résultats: %d artiste(s) trouvé(s)", len(matches))
		cards = append(cards, widget.NewCard("", "", widget.NewLabel(titleText)))

		for _, a := range matches {
			info := container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Nom: %s", a.Name)),
				widget.NewLabel(fmt.Sprintf("Membres: %v (%d)", a.Members, len(a.Members))),
				widget.NewLabel(fmt.Sprintf("Création: %d", a.CreationDate)),
				widget.NewLabel(fmt.Sprintf("Premier album: %s", a.FirstAlbum)),
			)

			locLabel := widget.NewLabel("Cliquer sur 'Voir lieux' pour charger les lieux")
			seeLoc := widget.NewButton("Voir lieux", func(aID int, ll *widget.Label) func() {
				return func() {
					rels, err := GetRelations()
					if err != nil {
						log.Printf("❌ Erreur API GetRelations pour artiste %d: %v\n", aID, err)
						ll.SetText("Erreur: " + err.Error())
						return
					}
					locations := GetArtistLocations(aID, rels)
					if len(locations) == 0 {
						log.Printf("⚠️  Aucun lieu trouvé pour l'artiste %d\n", aID)
						ll.SetText("Aucun lieu trouvé")
						return
					}
					log.Printf("✅ Lieux trouvés pour artiste %d: %d lieux\n", aID, len(locations))
					var sb strings.Builder
					sb.WriteString(fmt.Sprintf("Lieux (%d):\n", len(locations)))
					for _, loc := range locations {
						sb.WriteString(fmt.Sprintf("- %s\n", strings.Title(strings.ToLower(loc))))
					}
					ll.SetText(sb.String())
				}
			}(a.ID, locLabel))

			cardContent := container.NewVBox(info, locLabel, seeLoc)
			card := widget.NewCard(a.Name, "", cardContent)
			cards = append(cards, card)
		}

		resultContainer.Objects = cards
		resultContainer.Refresh()
		resultScroller.ScrollToTop()
	}

	// Fonction appelée lors des changements
	updateSearch := func() {
		if searchTimer != nil {
			searchTimer.Stop()
		}
		searchTimer = time.AfterFunc(300*time.Millisecond, func() {
			performSearch(searchEntry.Text, memberCountEntry.Text, creationYearEntry.Text)
		})
	}

	// Lier les changements des champs de recherche et filtres
	searchEntry.OnChanged = func(s string) {
		updateSearch()
	}
	memberCountEntry.OnChanged = func(s string) {
		updateSearch()
	}
	creationYearEntry.OnChanged = func(s string) {
		updateSearch()
	}

	// Bouton pour réinitialiser les filtres
	clearButton := widget.NewButton("Réinitialiser", func() {
		searchEntry.SetText("")
		memberCountEntry.SetText("")
		creationYearEntry.SetText("")
		resultContainer.Objects = []fyne.CanvasObject{
			widget.NewCard("", "", widget.NewLabel("Entrez des critères pour rechercher")),
		}
		resultContainer.Refresh()
	})

	// Interface des filtres
	filterBox := container.NewVBox(
		widget.NewLabel("=== Filtres ==="),
		searchEntry,
		memberCountEntry,
		creationYearEntry,
		clearButton,
	)

	filterScroller := container.NewVScroll(filterBox)
	filterScroller.SetMinSize(fyne.NewSize(300, 0))

	// Disposition principale: filtres à gauche, résultats à droite
	mainContent := container.NewBorder(
		nil, nil,
		filterScroller,
		nil,
		resultScroller,
	)

	window.SetContent(mainContent)
	window.ShowAndRun()
}
