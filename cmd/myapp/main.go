package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type PrayerTime struct {
	Data struct {
		Timings struct {
			Fajr    string `json:"Fajr"`
			Dhuhr   string `json:"Dhuhr"`
			Asr     string `json:"Asr"`
			Maghrib string `json:"Maghrib"`
			Isha    string `json:"Isha"`
		} `json:"timings"`
	} `json:"data"`
}

func main() {
	// Initialize the application
	myApp := app.New()
	myWindow := myApp.NewWindow("Prayer Times")

	// Create input fields for city and country
	cityEntry := widget.NewEntry()
	cityEntry.SetPlaceHolder("Enter City")

	countryEntry := widget.NewEntry()
	countryEntry.SetPlaceHolder("Enter Country")

	// Create labels for each prayer time
	fajrLabel := widget.NewLabel("Fajr: -")
	dhuhrLabel := widget.NewLabel("Dhuhr: -")
	asrLabel := widget.NewLabel("Asr: -")
	maghribLabel := widget.NewLabel("Maghrib: -")
	ishaLabel := widget.NewLabel("Isha: -")

	// Fetch and update prayer times
	updatePrayerTimes := func() {
		city := cityEntry.Text
		country := countryEntry.Text

		apiURL := fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity?city=%s&country=%s", city, country)
		resp, err := http.Get(apiURL)
		if err != nil {
			log.Printf("Error fetching prayer times: %v", err)
			return
		}
		defer resp.Body.Close()

		var prayerData PrayerTime
		if err := json.NewDecoder(resp.Body).Decode(&prayerData); err != nil {
			log.Printf("Failed to decode API response: %v", err)
			return
		}

		// Update labels with prayer times
		fajrLabel.SetText("Fajr: " + prayerData.Data.Timings.Fajr)
		dhuhrLabel.SetText("Dhuhr: " + prayerData.Data.Timings.Dhuhr)
		asrLabel.SetText("Asr: " + prayerData.Data.Timings.Asr)
		maghribLabel.SetText("Maghrib: " + prayerData.Data.Timings.Maghrib)
		ishaLabel.SetText("Isha: " + prayerData.Data.Timings.Isha)
	}

	// Button to get prayer times
	getTimesButton := widget.NewButton("Get Prayer Times", updatePrayerTimes)

	// Use a vertical box layout for the labels and inputs
	content := container.NewVBox(
		widget.NewLabel("Enter City and Country:"),
		cityEntry,
		countryEntry,
		getTimesButton,
		fajrLabel,
		dhuhrLabel,
		asrLabel,
		maghribLabel,
		ishaLabel,
	)

	// Set the content for the window and run the application
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300)) // Resize the window to a reasonable default
	myWindow.ShowAndRun()
}
