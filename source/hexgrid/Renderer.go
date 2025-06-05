package hexgrid

import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "math"

type Renderer struct {
	Canvas  *canvas2d.Canvas  `json:"canvas"`
	Context *canvas2d.Context `json:"context"`
	Grid    *Grid             `json:"grid"`
	Active  *Hexagon          `json:"active"`
	Hover   *Hexagon          `json:"hover"`
}

func NewRenderer(canvas *canvas2d.Canvas, grid *Grid) Renderer {

	var renderer Renderer

	renderer.Canvas = canvas
	renderer.Context = canvas.GetContext()
	renderer.Grid = grid
	renderer.Active = nil
	renderer.Hover = nil

	return renderer

}

func (renderer *Renderer) Clear() {
	renderer.Context.ClearRect(0, 0, int(renderer.Canvas.Width), int(renderer.Canvas.Height))
}

func (renderer *Renderer) Render() {

	renderer.Clear()

	renderer.Grid.ForEach(func(hexagon *Hexagon) {

		if renderer.Active != nil && renderer.Active == hexagon {
			// Do Nothing
		} else if renderer.Hover != nil && renderer.Hover == hexagon {
			// Do Nothing
		} else {
			renderer.RenderHexagon(hexagon)
		}

	})

	if renderer.Hover != nil {
		renderer.RenderHexagon(renderer.Hover)
	}

	if renderer.Active != nil {
		renderer.RenderHexagon(renderer.Active)
	}

}

func (renderer *Renderer) RenderEntity(entity *Entity) {

	context := renderer.Context

	pos_x, pos_y := entity.GetPosition()

	context.BeginPath()
	context.SetFillStyleColor("rgba(0,100,255,0.7)")
	context.Arc(pos_x, pos_y, 16, 0.0, 2 * math.Pi, false)
	context.Fill()
	context.ClosePath()


	// TODO: If Entity concept works, try to make an Entity for 2D effects
	// TODO: If Entity concept works, create also a GridEntity for hexagon layer
	// if entity.Hexagon != nil {

	// 	renderer.RenderHexagon(entity.Hexagon)

	// 	if entity.Image != nil {
	// 		// TODO: renderer.RenderImage(entity.Image, entity.Hexagon.Position)
	// 	}

	// 	if entity.Label != "" {
	// 		renderer.RenderLabel(entity.Label, entity.Hexagon.Position)
	// 	}

	// }

}

func (renderer *Renderer) RenderHexagon(hexagon *Hexagon) {

	context := renderer.Context
	polygon := renderer.Grid.ToScreenPolygon(*hexagon.Position)

	if renderer.Active != nil && renderer.Active.IsEqual(*hexagon) {

		context.BeginPath()
		context.SetFillStyleColor("rgba(0,100,255,0.7)")
		context.SetStrokeStyleColor("#ffffff")
		context.MoveTo(polygon[0].X, polygon[0].Y)
		context.LineTo(polygon[1].X, polygon[1].Y)
		context.LineTo(polygon[2].X, polygon[2].Y)
		context.LineTo(polygon[3].X, polygon[3].Y)
		context.LineTo(polygon[4].X, polygon[4].Y)
		context.LineTo(polygon[5].X, polygon[5].Y)
		context.LineTo(polygon[0].X, polygon[0].Y)
		context.Fill()
		context.Stroke()
		context.ClosePath()

	} else if renderer.Hover != nil && renderer.Hover.IsEqual(*hexagon) {

		context.BeginPath()
		context.SetFillStyleColor("rgba(255,255,255,0.5)")
		context.SetStrokeStyleColor("#ffffff")
		context.MoveTo(polygon[0].X, polygon[0].Y)
		context.LineTo(polygon[1].X, polygon[1].Y)
		context.LineTo(polygon[2].X, polygon[2].Y)
		context.LineTo(polygon[3].X, polygon[3].Y)
		context.LineTo(polygon[4].X, polygon[4].Y)
		context.LineTo(polygon[5].X, polygon[5].Y)
		context.LineTo(polygon[0].X, polygon[0].Y)
		context.Fill()
		context.Stroke()
		context.ClosePath()

	} else {

		context.BeginPath()
		context.SetFillStyleColor("rgba(255,255,255,0.1)")
		context.SetStrokeStyleColor("#ffffff")
		context.MoveTo(polygon[0].X, polygon[0].Y)
		context.LineTo(polygon[1].X, polygon[1].Y)
		context.LineTo(polygon[2].X, polygon[2].Y)
		context.LineTo(polygon[3].X, polygon[3].Y)
		context.LineTo(polygon[4].X, polygon[4].Y)
		context.LineTo(polygon[5].X, polygon[5].Y)
		context.LineTo(polygon[0].X, polygon[0].Y)
		context.Fill()
		context.Stroke()
		context.ClosePath()

	}

}

func (renderer *Renderer) SetHover(hexagon *Hexagon) {
	renderer.Hover = hexagon
}

func (renderer *Renderer) SetActive(hexagon *Hexagon) {
	renderer.Active = hexagon
}
