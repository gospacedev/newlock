package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/gocrazygh/newlock/assets"

	"github.com/sethvargo/go-password/password"
)

func main() {
	a := app.New()

	w := a.NewWindow("Newlock")

	//command to generate bundled.go
	//fyne bundle Icon.png > bundled.go
	w.SetIcon(assets.Icon)

	title := canvas.NewText("Newlock Password Generator", color.White)

	input := widget.NewEntry()

	input.SetPlaceHolder("Enter password length")

	var isUpperCase bool

	toggleUpperCase := widget.NewCheck("Upper Case Letters", func(value bool) {
		if !value {
			isUpperCase = true
		} else if value{
			isUpperCase = false
		}
	})

	toggleUpperCase.Checked = true

	text := canvas.NewText("", color.White)

	text.TextSize = 16

	btn := widget.NewButton("Generate", func() {
		passwordLength, _ := strconv.Atoi(input.Text)

		//Avoids length errors
		if passwordLength <= 10 || passwordLength >= 63 {
			passwordLength = 10
		}

		text.Text = password.MustGenerate(passwordLength, 5, 5, isUpperCase, false)

		text.Refresh()
	})

	copybtn := widget.NewButtonWithIcon("Copy Password", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(text.Text)
	})

	w.SetContent(container.NewVBox(
		title,
		input,
		toggleUpperCase,
		text,
		btn,
		copybtn,
	))

	w.Resize(fyne.NewSize(375, 300))

	w.ShowAndRun()

}
