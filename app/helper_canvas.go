package kss

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
Extend label widget to support being tapped.
Used to create rounded buttons.
*/
type TappableText struct {
	widget.Label

	OnTapped  func()
	Animation *fyne.Animation
}

func (tt *TappableText) Tapped(_ *fyne.PointEvent) {
	tt.OnTapped()
}

// Create rounded rectangle from corner radius,
// stroke width, background and stroke colors
func RoundedRectangle(radius, stroke float32, bgClr, strokeClr color.Color) (*fyne.Container, *fyne.Animation) {

	bg := canvas.NewRectangle(color.Transparent)
	bg.CornerRadius = radius
	bg.StrokeWidth = stroke
	bg.StrokeColor = strokeClr

	anim := canvas.NewColorRGBAAnimation(
		color.Transparent, bgClr, time.Millisecond*500, func(c color.Color) {
			bg.FillColor = c
			canvas.Refresh(bg)
		})
	anim.Start()

	return container.NewMax(bg), anim
}

// Create rounded rectangle with callback from corner radius,
// stroke width, background and stroke colors and a callback function
func RoundedButton(txt string, align fyne.TextAlign, radius, stroke float32, bgClr, strokeClr color.Color, callback func()) *fyne.Container {
	t := &TappableText{}
	t.Text = txt
	t.Alignment = align

	c, a := RoundedRectangle(radius, stroke, bgClr, strokeClr)
	c.Add(t)
	t.Animation = a

	t.OnTapped = func() {
		a.Start()
		callback()
	}
	return c
}

func MovingAnimation(obj fyne.CanvasObject, startPos, endPos fyne.Position, mlSec float32) fyne.CanvasObject {
	var anim *fyne.Animation

	duration := time.Millisecond * time.Duration(mlSec)

	anim = canvas.NewPositionAnimation(startPos, endPos, duration, func(pos fyne.Position) {
		obj.Move(pos)
		obj.Refresh()
	})
	anim.AutoReverse = true
	anim.Start()

	return obj

}

//

/*
Extend entry field to work around android bug
where entry is hidden behind a keyboard.
*/
type EntryField struct {
	widget.Entry

	isMobile   bool           // is mobile device
	SepCount   int            // number of separators (default: 35)
	Separators fyne.Container // container with separators (transparent lines)
}

func NewEntryField(win *fyne.Window, mobile bool) *EntryField {

	ef := &EntryField{
		isMobile:   mobile,
		Separators: *container.NewVBox(),
		SepCount:   35,
	}

	ef.ExtendBaseWidget(ef)

	if ef.isMobile {
		for x := 0; x < ef.SepCount; x++ {
			ef.Separators.Add(canvas.NewLine(color.Transparent))
		}
	}
	ef.Separators.Hide()
	return ef
}

func (ef *EntryField) TypedRune(r rune) {
	ef.Entry.TypedRune(r)
}

func (ef *EntryField) TypedKey(k *fyne.KeyEvent) {
	ef.Entry.TypedKey(k)
}

func (ef *EntryField) FocusGained() {
	ef.Entry.FocusGained()

	if ef.isMobile {
		ef.Separators.Show()
		ef.Refresh()
	}
}

func (ef *EntryField) FocusLost() {
	ef.Entry.FocusLost()

	if ef.isMobile {
		ef.Separators.Hide()
		ef.Refresh()
	}
}
