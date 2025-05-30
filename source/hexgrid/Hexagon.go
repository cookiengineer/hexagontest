package hexgrid

type Hexagon struct {
	Position *HexPosition `json:"position"`
}

func NewHexagon(q int, r int, s int) Hexagon {

	var hexagon Hexagon

	position := HexPosition{0,0,0}

	hexagon.Position = &position
	hexagon.SetPosition(q, r, s)

	return hexagon

}

func (hexagon *Hexagon) IsEqual(other Hexagon) bool {

	var result bool

	if hexagon.Position.Q == other.Position.Q && hexagon.Position.R == other.Position.R && hexagon.Position.S == other.Position.S {
		result = true
	}

	return result

}

func (hexagon *Hexagon) SetPosition(q int, r int, s int) bool {

	var result bool

	if q + r + s == 0 {

		hexagon.Position.Q = q
		hexagon.Position.R = r
		hexagon.Position.S = s

		result = true

	}

	return result

}
