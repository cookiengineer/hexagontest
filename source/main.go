package main

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/fetch"
import "battlemap/hexgrid"
import "battlemap/structs"
import "encoding/json"
import "math"
import "time"

func main() {

	element := dom.Document.QuerySelector("canvas")
	sidebar := dom.Document.QuerySelector("aside")
	canvas  := canvas2d.ToCanvas(element)

	hexagons := []hexgrid.Hexagon{
		hexgrid.NewHexagon( 0,  0,  0),
		hexgrid.NewHexagon( 1, -1,  0),
		hexgrid.NewHexagon(-1,  1,  0),
		hexgrid.NewHexagon( 0,  1, -1),
		hexgrid.NewHexagon( 0, -1,  1),
		hexgrid.NewHexagon( 1,  0, -1),
		hexgrid.NewHexagon(-1,  0,  1),
	}
	grid := hexgrid.NewMap(1024, 640, 64)

	for _, hexagon := range hexagons {
		grid.Add(&hexagon)
	}


	response, err1 := fetch.Fetch("http://localhost:3000/api/systems", &fetch.Request{
		Method: fetch.MethodGet,
		Mode:   fetch.ModeSameOrigin,
	})

	if err1 == nil {

		var system_names []string

		json.Unmarshal(response.Body, &system_names)

		for i, name := range system_names {

			response2, err2 := fetch.Fetch("http://localhost:3000/api/systems/" + name, &fetch.Request{
				Method: fetch.MethodGet,
				Mode:   fetch.ModeSameOrigin,
			})

			if err2 == nil {

				var system structs.System

				json.Unmarshal(response2.Body, &system)

				console.Log(system)

				systemHex := hexgrid.NewSystemHexagon(&system, i, 0, 0)
				grid.Add(&systemHex.Hexagon)  // Add the base Hexagon part to the grid

			}

		}
	}

	renderer := hexgrid.NewRenderer(canvas, &grid)

	canvas.Element.AddEventListener("mouseup", dom.ToEventListener(func(event *dom.Event) {

		bounding_rect   := canvas.Element.GetBoundingClientRect()
		screen_position := hexgrid.ScreenPosition{
			X: event.Value.Get("clientX").Int() - bounding_rect.X,
			Y: event.Value.Get("clientY").Int() - bounding_rect.Y,
		}

		grid_position := grid.ToHexPosition(screen_position)

		hexagon := grid.Get(grid_position.Q, grid_position.R, grid_position.S)

		if hexagon != nil {
			sidebar.SetAttribute("data-state", "active")
			renderer.SetHover(nil)
			renderer.SetActive(hexagon)
		} else {
			sidebar.RemoveAttribute("data-state")
			renderer.SetHover(nil)
			renderer.SetActive(nil)
		}

	}))

	var screen_position hexgrid.ScreenPosition

	canvas.Element.AddEventListener("mousemove", dom.ToEventListener(func(event *dom.Event) {

		bounding_rect   := canvas.Element.GetBoundingClientRect()
		screen_position = hexgrid.ScreenPosition{
			X: event.Value.Get("clientX").Int() - bounding_rect.X,
			Y: event.Value.Get("clientY").Int() - bounding_rect.Y,
		}

		grid_position := grid.ToHexPosition(screen_position)

		hexagon := grid.Get(grid_position.Q, grid_position.R, grid_position.S)

		if hexagon != nil {
			renderer.SetHover(hexagon)
		} else {
			renderer.SetHover(nil)
		}

	}))

	for true {

		animations.RequestAnimationFrame(func(timestamp float64) {

			renderer.Render()

			context := renderer.Context
			context.BeginPath()
			context.SetFillStyleColor("#ff0000")
			context.Arc(screen_position.X, screen_position.Y, 5, 0, 2.0 * math.Pi, false);
			context.Fill()
			context.ClosePath()

		})

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
