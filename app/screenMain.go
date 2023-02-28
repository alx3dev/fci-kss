package kss

import (
	theme2 "kss/theme"

	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fogleman/gg"
)

func (k *Kss) screenMain() fyne.CanvasObject {
	// Entries for data input.
	// (1)-edit to suit your needs.
	breedEntry := widget.NewEntry()
	nameEntry := widget.NewEntry()
	sexEntry := widget.NewEntry()
	birthEntry := widget.NewEntry()
	stoodEntry := widget.NewEntry()
	registerEntry := widget.NewEntry()
	tatooEntry := widget.NewEntry()
	ownerEntry := widget.NewEntry()
	addressEntry := widget.NewEntry()
	onEntry := widget.NewEntry()
	byEntry := widget.NewEntry()

	// Selectable data input for hip.
	// (2)-edit to suit your needs.
	hipSelect := widget.NewSelect([]string{"A", "B", "C", "D", "E"}, func(s string) {
		k.Hip = s
	})
	hipSelect.PlaceHolder = "HIP"
	hipSelect.Alignment = fyne.TextAlignCenter

	// Selectable data input for elbow.
	// (3)-edit to suit your needs.
	elbowSelect := widget.NewSelect([]string{"0", "BL", "1", "2", "3"}, func(s string) {
		k.Elbow = s
	})
	elbowSelect.PlaceHolder = "ELBOW"
	elbowSelect.Alignment = fyne.TextAlignCenter

	// Button to execute certificate creation.
	// PNG will be created in a same directory as executable.
	// (1a)-edit DrawCertificate to be same as (1).
	btnExec := RoundedButton("   Create Certificate   ", fyne.TextAlignCenter, 6, 1, theme.PrimaryColor(), theme.PrimaryColor(), func() {
		go func() {
			k.DrawCertificate(
				breedEntry.Text, nameEntry.Text, sexEntry.Text, birthEntry.Text, stoodEntry.Text,
				registerEntry.Text, tatooEntry.Text, ownerEntry.Text, addressEntry.Text, onEntry.Text, byEntry.Text,
			)

		}()
	})

	// Button to reset all fields.
	// (1b)-edit to be same as (1)
	btnReset := RoundedButton("   Reset Certificate   ", fyne.TextAlignCenter, 6, 1,
		color.RGBA{R: 90, G: 90, B: 90, A: 125}, color.RGBA{R: 90, G: 90, B: 90, A: 125}, func() {
			breedEntry.SetText("")
			nameEntry.SetText("")
			sexEntry.SetText("")
			birthEntry.SetText("")
			stoodEntry.SetText("")
			registerEntry.SetText("")
			tatooEntry.SetText("")
			ownerEntry.SetText("")
			addressEntry.SetText("")
			onEntry.SetText("")
			byEntry.SetText("")
			hipSelect.ClearSelected()
			elbowSelect.ClearSelected()
		})

	// Put buttons in a container with separators.
	btn := (container.NewVBox(
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
		container.NewCenter(container.NewHBox(btnExec, btnReset)),
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	))

	// Top form label.
	// (4)-edit to suit your needs.
	mainLabel := widget.NewButtonWithIcon("HIP-ELBOW DYSPLASIA - INTERNATIONAL CERTIFICATE", theme2.Icon, func() {})
	mainLabel.Alignment = widget.ButtonAlignCenter
	mainLabel.IconPlacement = widget.ButtonIconTrailingText

	// Create a top form with our entries.
	// (1c)-edit to be same as (1)
	formDog := widget.NewForm(
		widget.NewFormItem("Breed: ", breedEntry),
		widget.NewFormItem("Name: ", nameEntry),
		widget.NewFormItem("Sex: ", sexEntry),
		widget.NewFormItem("Birth: ", birthEntry),
		widget.NewFormItem("Stood Book: ", stoodEntry),
		widget.NewFormItem("Register: ", registerEntry),
		widget.NewFormItem("Chip No: ", tatooEntry),
		widget.NewFormItem("Owner: ", ownerEntry),
		widget.NewFormItem("Address: ", addressEntry),
	)

	// Bottom form label.
	// (5)-edit to suit your needs.
	byLabel := widget.NewLabel("THE EVALUATION WAS MADE:")
	byLabel.Alignment = fyne.TextAlignCenter

	// Create a bottom form for signature part.
	// (6)-edit to suit your needs.
	formEval :=
		widget.NewForm(
			widget.NewFormItem("On: ", onEntry),
			widget.NewFormItem("By: ", byEntry),
		)

	// Put forms and labels in a container,
	// with selectable widgets between them.
	// (7)-edit to suit your needs.
	form := container.NewVBox(
		mainLabel,
		formDog,
		widget.NewSeparator(), widget.NewSeparator(),
		container.NewAdaptiveGrid(2, hipSelect, elbowSelect),
		widget.NewSeparator(), widget.NewSeparator(),
		byLabel,
		formEval,
	)

	return container.NewBorder(form, btn, nil, nil, btn)
}

// Draw and save PNG certificate from arguments.
// (1d)-edit arguments to be same as (1)
// (1e)-edit dc.DrawString to be same as arguments, and edit your coordinates
func (k *Kss) DrawCertificate(breed, name, sex, birth, stood, register, tatoo, owner, address, on, by string) {
	// A4 in 150PPI
	dc := gg.NewContext(1754, 1240)

	rootUri := fyne.CurrentApp().Storage().RootURI().Path()

	dc.LoadFontFace(rootUri+"/FreeSansBoldOblique.ttf", 24)
	dc.SetLineWidth(2)
	dc.SetColor(color.Black)

	dc.DrawString(breed, 255, 399)
	dc.DrawString(name, 245, 428)
	dc.DrawString(sex, 235, 460)
	dc.DrawString(birth, 765, 460)
	dc.DrawString(stood, 1360, 460)
	dc.DrawString(register, 460, 550)
	dc.DrawString(tatoo, 1185, 550)
	dc.DrawString(owner, 300, 610)
	dc.DrawString(address, 310, 640)

	// (2a)-edit Y coordinates and radius for Hip rating
	if k.Hip != "" {
		dc.DrawCircle(hipFromSelect(k.Hip), 780, 23)
	}

	// (3a)-edit Y coordinates and radius for Elbow rating
	if k.Elbow != "" {
		dc.DrawCircle(elbowFromSelect(k.Elbow), 847, 23)
	}

	dc.Stroke()

	dc.DrawString(on, 220, 995)
	dc.DrawString(by, 905, 995)

	// Export in a Fyne storage directory.
	//ExportPNG("h.e.d"+"_name.png", dc.Image())

	// Export in a same directory where executable is.
	SavePNG(dc, name)
}

// Get coordinates for selected object (hip rating).
// (2b)-edit X coordinates for hip rating (2)
func hipFromSelect(s string) (x float64) {
	switch s {
	case "A":
		x = 898
	case "B":
		x = 975
	case "C":
		x = 1055
	case "D":
		x = 1130
	case "E":
		x = 1210
	}
	return x
}

// Get coordinates for selected object (elbow rating)
// (3b)-edit X coordinates for elbow rating (3)
func elbowFromSelect(s string) (x float64) {
	switch s {
	case "0":
		x = 894
	case "BL":
		x = 987
	case "1":
		x = 1055
	case "2":
		x = 1126
	case "3":
		x = 1195
	}
	return x
}

// Save PNG certificate in same folder where the app is.
// It doesn't check for symlinks, be careful
func SavePNG(ctx *gg.Context, name string) {
	ex, err := os.Executable()
	if err != nil {
		fyne.LogError("", err)
		return
	}
	path := filepath.Dir(ex)
	ctx.SavePNG(path + "/" + "h.e.d_" + name + ".png")
}
