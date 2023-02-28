package kss

import (
	"image"
	"image/png"
	theme2 "kss/theme"

	"time"

	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func (k *Kss) InitializeConfiguration() {
	k.handleFirstStart()
}

func (k *Kss) handleFirstStart() {
	fi := k.config.BoolWithFallback("FIRST_INIT", true)

	if fi {
		w := k.APP.NewWindow("Enter API Key")
		ent := widget.NewPasswordEntry()

		btn := widget.NewButton("OK", func() {
			if ent.Text != "1234567890" { // Why? Why not? :)
				k.APP.Quit()
			} else {
				k.config.SetBool("FIRST_INIT", false)
				ExportFromResource(theme2.TextFont)
				k.WIN.Show()
				w.Close()
			}
		})

		scr := container.NewVBox(
			widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
			ent,
			widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
			btn,
		)

		go func() {
			time.Sleep(time.Millisecond * 300)
			w.Resize(fyne.NewSize(k.width*1.8, k.height/1.8))
			w.CenterOnScreen()
			w.SetContent(scr)
			w.Show()
			k.WIN.Hide()
		}()
	}
}

func getOS() uint8 {
	switch runtime.GOOS {
	case "linux", "freebsd", "netbsd", "openbsd", "dragonfly":
		return 1
	case "android":
		return 2
	case "windows":
		return 3
	case "darwin":
		return 4
	case "ios":
		return 5
	}
	return 0
}

func (k *Kss) isMobile() bool {
	return k.APP.Driver().Device().IsMobile()
}

func ExportPNG(name string, img image.Image) {
	writer, _ := prepareWriter(name)

	err := png.Encode(writer, img)

	if err != nil {
		fyne.LogError("encoding image went wront", err)
	}
}

// Save fyne resource context to hard drive,
func ExportFromResource(res fyne.Resource) {
	Export(res.Name(), res.Content())
}

func Export(name string, data []byte) {
	writer, _ := prepareWriter(name)
	//os.WriteFile(name, data, 0600)
	writer.Write(data)
	writer.Close()
}

func prepareWriter(name string) (fyne.URIWriteCloser, error) {
	childURI, err := storage.Child(fyne.CurrentApp().Storage().RootURI(), name)
	if err != nil {
		fyne.LogError("app storage error", err)
	}

	writer, werr := storage.Writer(childURI)

	if werr != nil {
		fyne.LogError("check writing permissions", werr)
	}

	if err != nil {
		werr = err
	}

	return writer, werr
}
