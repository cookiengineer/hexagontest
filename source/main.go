package main

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/fetch"
import "battlemap/hexgrid"
import "battlemap/structs"
import "encoding/json"
import "time"

func main() {

	element := dom.Document.QuerySelector("canvas")
	sidebar := dom.Document.QuerySelector("aside")
	canvas  := canvas2d.ToCanvas(element)

	entity := hexgrid.Entity{
		Name:    "The Entity",
		Image:   nil,
		Effects: []*hexgrid.Effect{},
		System:  nil,
		Position: {
			X: 0,
			Y: 0,
		},
	}

	renderer := hexgrid.NewRenderer(canvas, nil)

	for true {

		now := time.Now()

		entity.Update(&now)

		animations.RequestAnimationFrame(func(timestamp float64) {
			renderer.RenderEntity(&entity)
		})

		time.Sleep(16 * time.Millisecond)

	}

}
