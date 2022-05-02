package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/sethvargo/go-password/password"
)


func main() {
	a := app.New()

	w := a.NewWindow("Newlock")

	r, _ := fyne.LoadResourceFromPath("assets/orange-lock.png")

	w.SetIcon(r)

	title := canvas.NewText("Newlock Password Generator", color.White)

	input := widget.NewEntry()

	input.SetPlaceHolder("Enter password length")

	text := canvas.NewText("", color.White)

	text.TextSize = 16

	btn := widget.NewButton("Generate", func() {
		//The number of available digits and symbols can be changes
		//support shorter passwords but it's still safe and random
		passwordLength,_ := strconv.Atoi(input.Text)
		
		if passwordLength == 0 {
		    passwordLength = 8
		}

		n := passwordLength / 2

		text.Text = password.MustGenerate(passwordLength, n, n, false, false)
	
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
