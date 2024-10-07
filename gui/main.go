package gui

import (
	"fmt"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	cram "github.com/aroario2003/cram/cmd"
)

func Show() {
	a := app.New()
	w := a.NewWindow("CRAM GUI")
	w.Resize(fyne.NewSize(1280, 720))

	//labels to display vulnerability score, time to fix, and top CVEs
	vulnScoreLabel := widget.NewLabel("Vulnerability Score: ")
	timeToFixLabel := widget.NewLabel("Time to Fix: ")
	topCVEsLabel := widget.NewLabel("Top CVEs: ")

	// input box
	queryInput := widget.NewEntry()
	queryInput.SetPlaceHolder("Enter query details here, seperated by , ...")

	conn := cram.ConnectToDbSocket()

	// buttons for each function
	queryDbOSBtn := widget.NewButton("QueryDbOS", func() {
		queryText := queryInput.Text
		cram.SetSoftware(queryText)
		result := cram.QueryDbOS(conn, queryText)
		rowsCount := cram.CountRowsReturned(result)
		vulnScoreLabel.SetText(fmt.Sprintf("Vulnerability Score: %f", cram.GetTotalVulnerabilityScore(result, rowsCount)))
		timeToFixLabel.SetText(fmt.Sprintf("Time to Fix: %d", cram.GetTotalTimeToFix(result)))
		topCVEsLabel.SetText(fmt.Sprintf("Top CVEs: %s", result))
	})

	queryDbCveBtn := widget.NewButton("QueryDbCve", func() {
		queryText := queryInput.Text
		cram.SetCveNum(queryText)
		result := cram.QueryDbCve(conn, queryText)
		rowsCount := cram.CountRowsReturned(result)
		vulnScoreLabel.SetText(fmt.Sprintf("Vulnerability Score: %f", cram.GetTotalVulnerabilityScore(result, rowsCount)))
		timeToFixLabel.SetText(fmt.Sprintf("Time to Fix: %d", cram.GetTotalTimeToFix(result)))
		topCVEsLabel.SetText(fmt.Sprintf("Top CVEs: %s", result))
	})

	queryDbMultiOsBtn := widget.NewButton("QueryDbMultiOs", func() {
		var vulnScoreStr string
		var ttfScoreStr string
		var topCveStr string

		queryText := queryInput.Text
		osList := strings.Split(queryText, ",")
		cram.SetSoftwares(osList)
		resultArr := cram.QueryDbMultiOs(osList)

		for i, result := range resultArr {
			// calculate vulnerability score and time to fix for each os provided
			rowsCount := cram.CountRowsReturned(result)
			vulnScore := cram.GetTotalVulnerabilityScore(result, rowsCount)
			ttfScore := cram.GetTotalTimeToFix(result)

			vulnScoreStr += fmt.Sprintf("%s: %f\n", osList[i], vulnScore) 
			ttfScoreStr += fmt.Sprintf("%s: %d\n", osList[i], ttfScore)
			topCveStr += fmt.Sprintf("%s: %s\n", osList[i], result)
		}

		vulnScoreLabel.SetText(fmt.Sprintf("Vulnerability Score(s): %s", vulnScoreStr))
		timeToFixLabel.SetText(fmt.Sprintf("Time(s) to Fix: %s", ttfScoreStr))
		topCVEsLabel.SetText(fmt.Sprintf("Top CVEs: %s", topCveStr))
	})

	queryDbMultiCveBtn := widget.NewButton("QueryDbMultiCve", func() {
		var vulnScoreStr string
		var ttfScoreStr string
		var topCveStr string

		queryText := queryInput.Text
		cveList := strings.Split(queryText, ",")
		cram.SetCveNums(cveList)
		resultArr := cram.QueryDbMultiCve(cveList)

		for i, result := range resultArr {
			// calculate vulnerability score and time to fix for each cve provided
			rowsCount := cram.CountRowsReturned(result)
			vulnScore := cram.GetTotalVulnerabilityScore(result, rowsCount)
			ttfScore := cram.GetTotalTimeToFix(result)

			vulnScoreStr += fmt.Sprintf("%s: %f\n", cveList[i], vulnScore) 
			ttfScoreStr += fmt.Sprintf("%s: %d\n", cveList[i], ttfScore)
			topCveStr += fmt.Sprintf("%s: %s\n", cveList[i], result)
		}

		vulnScoreLabel.SetText(fmt.Sprintf("Vulnerability Score(s): %s", vulnScoreStr))
		timeToFixLabel.SetText(fmt.Sprintf("Time(s) to Fix: %s", ttfScoreStr))
		topCVEsLabel.SetText(fmt.Sprintf("Top CVEs: %s", topCveStr))
	})

	// UI
	buttons := container.NewHBox(queryDbOSBtn, queryDbCveBtn, queryDbMultiOsBtn, queryDbMultiCveBtn)
	querySection := container.NewVBox(queryInput, buttons)
	leftPane := container.NewVBox(vulnScoreLabel)
	rightPane := container.NewVBox(timeToFixLabel)
	bottomPane := container.NewVBox(topCVEsLabel)

	// combining sections
	layout := container.NewBorder(querySection, bottomPane, leftPane, rightPane, nil)
	w.SetContent(layout)

	w.ShowAndRun()
}
