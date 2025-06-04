package hexgrid

import "battlemap/structs"

type Entity struct {
	Name    string          `json:"name"`
	Image   *Image          `json:"image"`
	Hexagon *Hexagon        `json:"hexagon"`
	System  *structs.System `json:"system"`
}

// TODO: How to represent entity.Image?
// Find out if go's internal image package makes sense
