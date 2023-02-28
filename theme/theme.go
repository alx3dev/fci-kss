package theme

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct {
	Theme string
}

//go:embed Icon.png
var Ico []byte
var Icon = &fyne.StaticResource{
	StaticName:    "Icon.png",
	StaticContent: Ico,
}

/*
Font to be used in app
*/
//go:embed OpenSans-Regular.ttf
var font []byte //(looking for better fonts...)
var MyFont = &fyne.StaticResource{
	StaticName:    "OpenSans-Regular.ttf",
	StaticContent: font,
}

/*
Font to be downloaded on first start,
and used for certificate
*/
//go:embed "FreeSansBoldOblique.ttf"
var textFont []byte
var TextFont = &fyne.StaticResource{
	StaticName:    "FreeSansBoldOblique.ttf",
	StaticContent: textFont,
}

var _ fyne.Theme = (*MyTheme)(nil)

func (m *MyTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return MyFont
}

func (m *MyTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {

	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 0
	case theme.SizeNameScrollBarSmall:
		return 0
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 1
	}

	return theme.DefaultTheme().Size(n)
}

func (m *MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch m.Theme {
	case "Dark":
		v = theme.VariantDark

	case "Light":
		v = theme.VariantLight
	}

	switch n {
	case theme.ColorNameSeparator:
		return color.Transparent

	case theme.ColorNamePrimary:
		return theme.WarningColor()

	case theme.ColorNameForeground:
		if v == theme.VariantLight {
			return color.RGBA{R: 30, G: 30, B: 30, A: 255}
		}

	case theme.ColorNameInputBackground:
		return color.Transparent
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
