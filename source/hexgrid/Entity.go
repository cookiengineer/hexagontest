package hexgrid

import (
	"battlemap/structs"
	"github.com/cookiengineer/gooey/bindings/canvas2d"
)

type SystemHexagon struct {
	Hexagon
	System *structs.System
}

func NewSystemHexagon(system *structs.System, q int, r int, s int) SystemHexagon {
	baseHexagon := NewHexagon(q, r, s)
	systemHex := SystemHexagon{
		Hexagon: baseHexagon,
		System:  system,
	}

	return systemHex
}

// render is now a private method that will be called through the field
func (sh *SystemHexagon) render(context *canvas2d.Context, hexMap *Map) {
	// Get the screen polygon points for this hexagon
	polygon := hexMap.ToScreenPolygon(*sh.Position)

	// Set styles for the hexagon
	context.BeginPath()
	context.SetFillStyleColor("#4a90e2")   // A nice blue color
	context.SetStrokeStyleColor("#2171c7") // Darker blue for border

	// Draw the hexagon path
	context.MoveTo(polygon[0].X, polygon[0].Y)
	for i := 1; i < len(polygon); i++ {
		context.LineTo(polygon[i].X, polygon[i].Y)
	}
	context.LineTo(polygon[0].X, polygon[0].Y)

	context.Fill()
	context.Stroke()
	context.ClosePath()

	// Calculate center point for text (average of all points)
	var sumX int
	var sumY int
	for _, point := range polygon {
		sumX += point.X
		sumY += point.Y
	}
	centerX := sumX / len(polygon)
	centerY := sumY / len(polygon)

	// Draw the system name
	context.BeginPath()
	context.SetFillStyleColor("#ffffff") // White text
	context.SetFont("14px Arial")
	context.SetTextAlign("center")
	context.SetTextBaseline("middle")
	context.FillText(sh.System.Hostname, centerX, centerY)
	context.ClosePath()
}
