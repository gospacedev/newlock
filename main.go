package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func main() {
	a := app.New()

	w := a.NewWindow("Password Generator")

	title := canvas.NewText("Simple Password Generator", color.White)

	input := widget.NewEntry()

	input.SetPlaceHolder("Enter password length")

	text := canvas.NewText("", color.White)

	text.TextSize = 16

	btn := widget.NewButton("Geneate a password", func() {
		passwordLength, _ := strconv.Atoi(input.Text)

		text.Text = PasswordGenerator(passwordLength)
	
		text.Refresh()
	})

	w.SetContent(container.NewVBox(
		title,
		input,
		text,
		btn,
	))

	w.Resize(fyne.NewSize(375, 300))

	w.ShowAndRun()

}