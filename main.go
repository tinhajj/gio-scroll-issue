package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"scroll/custom"
	"scroll/customWidget"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

type Component struct {
	list       customWidget.List
	clickables []widget.Clickable
}

func (c *Component) Layout(theme *material.Theme, gtx C) D {
	return custom.List(theme, &c.list).Layout(gtx, len(c.clickables), func(gtx C, i int) D {
		clickable := &c.clickables[i]
		var label string
		label = "FRONT Long Text Button !!!!!!!!!!! END"
		if i%2 == 0 {
			label = "FRONT Long Text Button END"
		}
		return material.Button(theme, clickable, label).Layout(gtx)
	})
}

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Canvas"),
			app.Size(unit.Dp(200), unit.Dp(350)),
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
		clickables: []widget.Clickable{{}, {}, {}, {}, {}, {}},
	}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			gtx.Constraints.Min = image.Pt(0, 0)

			for _, e := range gtx.Events(3) {
				fmt.Println(e)
			}

			mycomp.Layout(theme, gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
