package hexgrid

import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "github.com/cookiengineer/gooey/bindings/console"

type Renderer struct {
	Canvas  *canvas2d.Canvas  `json:"canvas"`
	Context *canvas2d.Context `json:"context"`
	Map     *Map              `json:"map"`
	Active  *Hexagon          `json:"active"`
	Palette struct {
		Default string `json:"default"`
		Active  string `json:"active"`
	} `json:"palette"`
}

func NewRenderer(canvas *canvas2d.Canvas, themap *Map) Renderer {

	var renderer Renderer

	renderer.Canvas = canvas
	renderer.Context = canvas.GetContext()
	renderer.Map = themap
	renderer.Active = nil
	renderer.Palette.Default = "#444444"
	renderer.Palette.Active = "#999999"

	return renderer

}

func (renderer *Renderer) Render() {

	renderer.Map.Each(func(hexagon *Hexagon) {

		context := renderer.Context
		polygon := renderer.Map.ToScreenPolygon(*hexagon.Position)

		if renderer.Active != nil {

			if renderer.Active.IsEqual(*hexagon) {

				context.BeginPath()
				context.SetFillStyleColor(renderer.Palette.Active)
				context.MoveTo(polygon[0].X, polygon[0].Y)
				context.LineTo(polygon[1].X, polygon[1].Y)
				context.LineTo(polygon[2].X, polygon[2].Y)
				context.LineTo(polygon[3].X, polygon[3].Y)
				context.LineTo(polygon[4].X, polygon[4].Y)
				context.LineTo(polygon[5].X, polygon[5].Y)
				context.LineTo(polygon[0].X, polygon[0].Y)
				context.Fill()
				context.ClosePath()

			} else {

				context.BeginPath()
				context.SetStrokeStyleColor(renderer.Palette.Default)
				context.MoveTo(polygon[0].X, polygon[0].Y)
				context.LineTo(polygon[1].X, polygon[1].Y)
				context.LineTo(polygon[2].X, polygon[2].Y)
				context.LineTo(polygon[3].X, polygon[3].Y)
				context.LineTo(polygon[4].X, polygon[4].Y)
				context.LineTo(polygon[5].X, polygon[5].Y)
				context.LineTo(polygon[0].X, polygon[0].Y)
				context.Stroke()
				context.ClosePath()

			}

		} else {

			context.BeginPath()
			context.SetStrokeStyleColor(renderer.Palette.Default)
			context.MoveTo(polygon[0].X, polygon[0].Y)
			context.LineTo(polygon[1].X, polygon[1].Y)
			context.LineTo(polygon[2].X, polygon[2].Y)
			context.LineTo(polygon[3].X, polygon[3].Y)
			context.LineTo(polygon[4].X, polygon[4].Y)
			context.LineTo(polygon[5].X, polygon[5].Y)
			context.LineTo(polygon[0].X, polygon[0].Y)
			context.Stroke()
			context.ClosePath()

		}

		console.Log(hexagon)
		console.Log(polygon)

	})

}

func (renderer *Renderer) SetActive(hexagon *Hexagon) {
	renderer.Active = hexagon
}
