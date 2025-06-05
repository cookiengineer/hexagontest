package main

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "github.com/cookiengineer/gooey/bindings/dom"
import "battlemap/hexgrid"
import "time"

func main() {

	element := dom.Document.QuerySelector("canvas")
	canvas  := canvas2d.ToCanvas(element)

	entity := hexgrid.Entity{
		Name:   "The Entity",
		Image:  nil,
		System: nil,
	}
	entity.SetPosition(0, 0)
	entity.SetRotation(0.0)
	entity.SetScale(1.0)

	renderer := hexgrid.NewRenderer(canvas, nil)

	go func() {

		time.Sleep(1 * time.Second)

		entity.TweenPosition(100, 100)

	}()

	for true {

		now := time.Now()
		entity.Update(&now)

		animations.RequestAnimationFrame(func(timestamp float64) {
			renderer.Clear()
			renderer.RenderEntity(&entity)
		})

		time.Sleep(16 * time.Millisecond)

	}

}
