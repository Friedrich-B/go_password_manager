package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	testApp := app.New()
	testWindow := testApp.NewWindow("Hello fellow kids!")
	testLabel := widget.NewLabel("asdf")
	testButton := widget.NewButton(
		"test",
		func() {
			testLabel.SetText("omg")
		},
	)
	testBox := container.NewVBox(testLabel, testButton)

	testWindow.SetContent(testBox)
	testWindow.ShowAndRun()
}
