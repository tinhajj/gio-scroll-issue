package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

type Component struct {
	list       widget.List
	clickables []widget.Clickable
}

func (c Component) Layout(theme *material.Theme, gtx C) D {
	return material.List(theme, &c.list).Layout(gtx, 3, func(gtx C, i int) D {
		clickable := &c.clickables[i]
		return material.Button(theme, clickable, "Long Text Button !!!!!!!").Layout(gtx)
	})
}

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Canvas"),
			app.Size(unit.Dp(200), unit.Dp(50)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops

	var theme = material.NewTheme(gofont.Collection())
	var mycomp = &Component{
		list: widget.List{
			Scrollbar: widget.Scrollbar{},
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		clickables: []widget.Clickable{{}, {}, {}},
	}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			gtx.Constraints.Min = image.Pt(0, 0)

			dim := mycomp.Layout(theme, gtx)
			rect := clip.Rect{
				Max: dim.Size,
			}
			s := rect.Push(gtx.Ops)
			paint.Fill(gtx.Ops, color.NRGBA{R: 255, A: 50})
			s.Pop()

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
