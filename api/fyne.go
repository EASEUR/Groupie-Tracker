package api

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FyneWindow() {
	appNew := app.New()
	window := appNew.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(800, 600))

	// Barre de recherche
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	// Zone de résultats
	resultLabel := widget.NewLabel("Les résultats apparaîtront ici")
	resultContainer := container.NewVBox(
		widget.NewCard("Résultats", "", resultLabel),
	)

	// Bouton de recherche
	searchButton := widget.NewButton("Rechercher", func() {
		query := searchEntry.Text
		if query == "" {
			resultLabel.SetText("Veuillez entrer un terme de recherche")
			return
		}
		resultLabel.SetText("Résultats pour: " + query)
	})

	// Layout principal
	searchBox := container.NewBorder(
		nil, nil, nil, searchButton,
		searchEntry,
	)

	mainContent := container.NewBorder(
		searchBox, nil, nil, nil,
		resultContainer,
	)

	window.SetContent(mainContent)
	window.ShowAndRun()
}
