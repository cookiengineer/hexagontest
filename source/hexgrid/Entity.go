package hexgrid

import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "battlemap/structs"
import "time"

type Entity struct {
	Name    string          `json:"name"`
	Image   *canvas2d.Image `json:"image"`
	Effects []Effect        `json:"effects"`
	System  *structs.System `json:"system"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

func (entity *Entity) Update(now *time.Time) {

	keep := make([]Effect, 0)

	for _, effect := range entity.Effects {

		result := effect.Update(now)

		if result == true {
			keep = append(keep, effect)
		}

	}

	entity.Effects = keep

}

func (entity *Entity) GetPosition() (int, int) {
	return entity.Position.X, entity.Position.Y
}

func (entity *Entity) SetPosition(x int, y int) {
	entity.Position.X = x
	entity.Position.Y = y
}

func (entity *Entity) ToPosition(x int, y int) {

	effect := PositionEffect{
		Entity:   entity,
		Start:    nil,
		Duration: 200 * time.Millisecond,
		From:     [2]int{},
		To:       [2]int{x, y},
	}

	entity.Effects = append(entity.Effects, effect)

}

