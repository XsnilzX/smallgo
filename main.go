package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	playBtn      widget.Clickable
	pauseBtn     widget.Clickable
	skipForward  widget.Clickable
	skipBackward widget.Clickable
	imageBtn     widget.Clickable
	progress     float32 = 0.5 // Dummy-Wert: 50% Fortschritt
)

func main() {
	go func() {
		window := new(app.Window)
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,

				// Box 1
				layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
					return layoutBox(gtx, theme, "Box 1", color.NRGBA{R: 200, G: 100, B: 100, A: 255})
				}),

				// Box 2
				layout.Flexed(0.7, func(gtx layout.Context) layout.Dimensions {
					return layoutBox(gtx, theme, "Box 2", color.NRGBA{R: 100, G: 200, B: 100, A: 255})
				}),

				// Box 3 mit Player Navigation
				layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
					return playerBox(gtx, theme)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func layoutBox(gtx layout.Context, theme *material.Theme, text string, bgColor color.NRGBA) layout.Dimensions {
	size := gtx.Constraints.Max

	// Hintergrund
	paint.FillShape(gtx.Ops,
		bgColor,
		clip.Rect{Max: size}.Op(),
	)

	// Label zentriert
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		lbl := material.H5(theme, text)
		lbl.Color = color.NRGBA{A: 255}
		return lbl.Layout(gtx)
	})
}

// playerBox zeichnet die Player-Steuerung in Box 3
func playerBox(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	size := gtx.Constraints.Max

	// Hintergrund
	paint.FillShape(gtx.Ops,
		color.NRGBA{R: 100, G: 100, B: 200, A: 255},
		clip.Rect{Max: size}.Op(),
	)

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,

		// Player Buttons
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Spacing: layout.SpaceAround,
				}.Layout(gtx,

					// Bild Button
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &imageBtn, "🖼️")
						return btn.Layout(gtx)
					}),

					// Skip Backward
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &skipBackward, "⏮️")
						return btn.Layout(gtx)
					}),

					// Play
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &playBtn, "▶️")
						return btn.Layout(gtx)
					}),

					// Pause
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &pauseBtn, "⏸️")
						return btn.Layout(gtx)
					}),

					// Skip Forward
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &skipForward, "⏭️")
						return btn.Layout(gtx)
					}),
				)
			})
		}),

		// Progress Bar unten
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			progressBar := material.ProgressBar(theme, progress)
			progressBar.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
			progressBar.TrackColor = color.NRGBA{A: 100}
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, progressBar.Layout)
		}),
	)
}
