package hexgrid

import "math"

type orientation struct {
	F0 float64
	F1 float64
	F2 float64
	F3 float64
	B0 float64
	B1 float64
	B2 float64
	B3 float64
	StartAngle float64
}

var orientation_flat orientation = orientation{
	F0: 3.0 / 2.0,
	F1: 0.0,
	F2: math.Sqrt(3.0) / 2.0,
	F3: math.Sqrt(3.0),
	B0: 2.0 / 3.0,
	B1: 0.0,
	B2: -1.0 / 3.0,
	B3: math.Sqrt(3.0) / 3.0,
	StartAngle: 0.0,
}

var orientation_pointy orientation = orientation{
	F0: math.Sqrt(3.0),
	F1: math.Sqrt(3.0) / 2.0,
	F2: 0.0,
	F3: 3.0 / 2.0,
	B0: math.Sqrt(3.0) / 3.0,
	B1: -1.0 / 3.0,
	B2: 0.0,
	B3: 2.0 / 3.0,
	StartAngle: 0.5,
}

