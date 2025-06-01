package hexgrid

import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "math"

type Renderer struct {
	Canvas  *canvas2d.Canvas  `json:"canvas"`
	Context *canvas2d.Context `json:"context"`
	Map     *Map              `json:"map"`
	Active  *Hexagon          `json:"active"`
	Hover   *Hexagon          `json:"hover"`
}

func NewRenderer(canvas *canvas2d.Canvas, themap *Map) Renderer {

	var renderer Renderer

	renderer.Canvas = canvas
	renderer.Context = canvas.GetContext()
	renderer.Map = themap
	renderer.Active = nil
	renderer.Hover = nil

	return renderer

}

func (renderer *Renderer) Render() {

	// Use default rendering (existing code)
	context := renderer.Context

	context.ClearRect(0, 0, int(renderer.Canvas.Width), int(renderer.Canvas.Height))

	context.BeginPath()
	context.SetFillStyleColor("#00ff00")
	context.Arc(renderer.Map.Origin.X, renderer.Map.Origin.Y, 5, 0, 2.0 * math.Pi, false);
	context.Fill()
	context.ClosePath()

	renderer.Map.Each(func(hexagon *Hexagon) {

		polygon := renderer.Map.ToScreenPolygon(*hexagon.Position)
		point   := renderer.Map.ToScreenPosition(*hexagon.Position)

		context.BeginPath()
		context.SetStrokeStyleColor("#ff0000")
		context.Arc(point.X, point.Y, 5, 0, 2.0 * math.Pi, false);
		context.Stroke()
		context.ClosePath()

		if renderer.Active != nil {

			if renderer.Active.IsEqual(*hexagon) {

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

		} else if renderer.Hover != nil {

			if renderer.Hover.IsEqual(*hexagon) {

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

	})

}

func (renderer *Renderer) SetHover(hexagon *Hexagon) {
	renderer.Hover = hexagon
}

func (renderer *Renderer) SetActive(hexagon *Hexagon) {
	renderer.Active = hexagon
}
