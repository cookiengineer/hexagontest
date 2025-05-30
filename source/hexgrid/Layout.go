package hexgrid

type Layout string

const (
	LayoutFlat   Layout = "flat"   // rotated 30deg
	LayoutPointy Layout = "pointy" // rotated 0deg (pointy top hexagons)
)
