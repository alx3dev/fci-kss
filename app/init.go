package kss

import (
	theme2 "kss/theme"

	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type Kss struct {
	WIN    fyne.Window      // Master window
	APP    fyne.App         // Fyne app
	config fyne.Preferences // App preferences

	system uint8 // linux-1, android-2, win-3, mac-4, iOS-5, unknown-0

	width  float32 // Master window width
	height float32 // Master window height

	Hip, Elbow string // Track selected options
}

func Initialize(id string) *Kss {
	myApp := app.NewWithID(id)

	myApp.Settings().SetTheme(&theme2.MyTheme{Theme: "Light"})

	kssApp := &Kss{
		APP:    myApp,
		WIN:    myApp.NewWindow(""),
		config: myApp.Preferences(),
		system: getOS(),
	}

	kssApp.InitializeConfiguration()

	kssApp.WIN.SetTitle("КСС Сертификат")

	return kssApp
}

// Draw screens and run application
func (k *Kss) Start() {

	scr := container.NewPadded(k.screenMain())

	k.WIN.SetContent(scr)

	k.width = k.WIN.Canvas().Size().Width
	k.height = k.WIN.Canvas().Size().Height

	if !k.isMobile() {
		scale := k.config.StringWithFallback("SCALE", "1")
		os.Setenv("FYNE_SCALE", scale)

		k.WIN.Resize(fyne.NewSize(k.width*2, k.height))
	}

	k.WIN.CenterOnScreen()
	k.WIN.SetMaster()
	k.WIN.Show()

	k.APP.Run()
}
