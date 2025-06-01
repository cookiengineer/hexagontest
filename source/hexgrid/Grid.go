package hexgrid

import "math"

type Grid struct {
	Layout Layout                           `json:"layout"`
	Origin *ScreenPosition                  `json:"center"`
	Width  int                              `json:"width"`
	Height int                              `json:"height"`
	Size   int                              `json:"size"`
	Map    map[int]map[int]map[int]*Hexagon `json:"grid"`
}

func NewGrid(width int, height int, size int) Grid {

	var grid Grid

	origin := ScreenPosition{width / 2, height / 2}

	grid.Layout = LayoutFlat
	grid.Origin = &origin
	grid.Width = width
	grid.Height = height
	grid.Size = size
	grid.Map = make(map[int]map[int]map[int]*Hexagon)

	return grid

}

func (grid *Grid) Add(hexagon *Hexagon) bool {

	var result bool

	q := hexagon.Position.Q
	r := hexagon.Position.R
	s := hexagon.Position.S

	_, ok1 := grid.Map[q]

	if ok1 == false {
		grid.Map[q] = make(map[int]map[int]*Hexagon)
		ok1 = true
	}

	_, ok2 := grid.Map[q][r]

	if ok2 == false {
		grid.Map[q][r] = make(map[int]*Hexagon)
		ok2 = true
	}

	_, ok3 := grid.Map[q][r][s]

	if ok3 == false {
		grid.Map[q][r][s] = hexagon
		result = true
	}

	return result

}

type Callback func(*Hexagon)

func (grid *Grid) ForEach(callback Callback) {

	min_q := int(-1 * float64(grid.Width) / float64(grid.Size))
	max_q := int( 1 * float64(grid.Width) / float64(grid.Size))

	// This is somewhat the actual distance, but angle changes based on layout
	// distance := math.Sqrt(math.Pow(float64(width) / 2, 2) + math.Pow(float64(height) / 2))

	min_r := int(-1 * (float64(grid.Height) / 2) / float64(grid.Size))
	max_r := int( 1 * (float64(grid.Height) / 2) / float64(grid.Size))

	min_s := int(-1 * (float64(grid.Width) / 2) / float64(grid.Size))
	max_s := int( 1 * (float64(grid.Width) / 2) / float64(grid.Size))

	for q := min_q; q <= max_q; q++ {

		for r := min_r; r <= max_r; r++ {

			for s := min_s; s <= max_s; s++ {

				tmp := grid.Get(q, r, s)

				if tmp != nil {
					callback(tmp)
				}

			}

		}

	}

}

func (grid *Grid) Get(q int, r int, s int) *Hexagon {

	var result *Hexagon = nil

	_, ok1 := grid.Map[q]

	if ok1 == true {

		_, ok2 := grid.Map[q][r]

		if ok2 == true {

			tmp, ok3 := grid.Map[q][r][s]

			if ok3 == true {
				result = tmp
			}

		}

	}

	return result

}

func (grid *Grid) Remove(hexagon *Hexagon) bool {

	var result bool

	q := hexagon.Position.Q
	r := hexagon.Position.R
	s := hexagon.Position.S

	_, ok1 := grid.Map[q]

	if ok1 == true {

		_, ok2 := grid.Map[q][r]

		if ok2 == true {

			_, ok3 := grid.Map[q][r][s]

			if ok3 == true {
				delete(grid.Map[q][r], s)
				result = true
			}

		}

	}

	return result

}

func (grid *Grid) SetOrigin(position ScreenPosition) bool {

	var result bool

	grid.Origin = &position

	return result

}

func (grid *Grid) ToScreenPosition(position HexPosition) ScreenPosition {

	var translated ScreenPosition

	if grid.Layout == LayoutFlat {

		orientation := orientation_flat
		x := float64(orientation.F0 * float64(position.Q) + orientation.F1 * float64(position.R)) * float64(grid.Size)
		y := float64(orientation.F2 * float64(position.Q) + orientation.F3 * float64(position.R)) * float64(grid.Size)

		translated.X = grid.Origin.X + int(x)
		translated.Y = grid.Origin.Y + int(y)

	} else if grid.Layout == LayoutPointy {

		orientation := orientation_pointy
		x := float64(orientation.F0 * float64(position.Q) + orientation.F1 * float64(position.R)) * float64(grid.Size)
		y := float64(orientation.F2 * float64(position.Q) + orientation.F3 * float64(position.R)) * float64(grid.Size)

		translated.X = grid.Origin.X + int(x)
		translated.Y = grid.Origin.Y + int(y)

	}

	return translated

}

func (grid *Grid) ToScreenPolygon(position HexPosition) []ScreenPosition {

	result := make([]ScreenPosition, 6)
	center := grid.ToScreenPosition(position)

	if grid.Layout == LayoutFlat {

		orientation := orientation_flat

		for corner := 0; corner < 6; corner++ {

			angle := float64(2.0 * math.Pi * (orientation.StartAngle + float64(corner)) / 6.0)

			result[corner] = ScreenPosition{
				X: int(float64(center.X) + float64(float64(grid.Size) * math.Cos(angle))),
				Y: int(float64(center.Y) + float64(float64(grid.Size) * math.Sin(angle))),
			}

		}

	} else if grid.Layout == LayoutPointy {

		orientation :=  orientation_pointy

		for corner := 0; corner < 6; corner++ {

			angle := float64(2.0 * math.Pi * (orientation.StartAngle + float64(corner)) / 6.0)

			result[corner] = ScreenPosition{
				X: int(float64(center.X) + float64(float64(grid.Size) * math.Cos(angle))),
				Y: int(float64(center.Y) + float64(float64(grid.Size) * math.Sin(angle))),
			}

		}

	}

	return result

}

func (grid *Grid) ToHexPosition(position ScreenPosition) HexPosition {

	var translated HexPosition

	tile_x := (float64(position.X) - float64(grid.Origin.X)) / float64(grid.Size)
	tile_y := (float64(position.Y) - float64(grid.Origin.Y)) / float64(grid.Size)

	if grid.Layout == LayoutFlat {

		orientation := orientation_flat

		q := float64(orientation.B0 * tile_x + orientation.B1 * tile_y)
		r := float64(orientation.B2 * tile_x + orientation.B3 * tile_y)
		s := float64(-1 * q - r)

		q_rounded := math.Round(q)
		r_rounded := math.Round(r)
		s_rounded := math.Round(s)

		q_diff := math.Abs(q_rounded - q)
		r_diff := math.Abs(r_rounded - r)
		s_diff := math.Abs(s_rounded - s)

		if q_diff > r_diff && q_diff > s_diff {

			translated.Q = int(-1 * r_rounded - s_rounded)
			translated.R = int(r_rounded)
			translated.S = int(s_rounded)

		} else if r_diff > s_diff {

			translated.Q = int(q_rounded)
			translated.R = int(-1 * q_rounded - s_rounded)
			translated.S = int(s_rounded)

		} else {

			translated.Q = int(q_rounded)
			translated.R = int(r_rounded)
			translated.S = int(-1 * q_rounded - r_rounded)

		}

	} else if grid.Layout == LayoutPointy {

		orientation := orientation_pointy

		q := float64(orientation.B0 * tile_x + orientation.B1 * tile_y)
		r := float64(orientation.B2 * tile_x + orientation.B3 * tile_y)
		s := float64(-1 * q - r)

		q_rounded := math.Round(q)
		r_rounded := math.Round(r)
		s_rounded := math.Round(s)

		q_diff := math.Abs(q_rounded - q)
		r_diff := math.Abs(r_rounded - r)
		s_diff := math.Abs(s_rounded - s)

		if q_diff > r_diff && q_diff > s_diff {

			translated.Q = int(-1 * r_rounded - s_rounded)
			translated.R = int(r_rounded)
			translated.S = int(s_rounded)

		} else if r_diff > s_diff {

			translated.Q = int(q_rounded)
			translated.R = int(-1 * q_rounded - s_rounded)
			translated.S = int(s_rounded)

		} else {

			translated.Q = int(q_rounded)
			translated.R = int(r_rounded)
			translated.S = int(-1 * q_rounded - r_rounded)

		}

	}

	return translated

}
