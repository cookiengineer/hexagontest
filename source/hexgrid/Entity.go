package hexgrid

import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "battlemap/structs"
import "battlemap/hexgrid/effects"
import "math"
import "time"

type Entity struct {
	Name    string          `json:"name"`
	Image   *canvas2d.Image `json:"image"`
	Effects struct {
		Position *effects.Position `json:"position"`
		Rotation *effects.Rotation `json:"rotation"`
		Scale    *effects.Scale    `json:"scale"`
	} `json:"effects"`
	System  *structs.System `json:"system"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	Rotation float64 `json:"rotation"`
	Scale    float64 `json:"scale"`
}

func NewEntity(name string, url string, x int, y int) *Entity {

	var entity Entity

	// TODO: System parameter

	if url != "" {
		image := canvas2d.NewImage(64, 64, url)
		entity.Image = &image
	} else {
		entity.Image = nil
	}

	entity.Name = name
	entity.Effects.Position = nil
	entity.Effects.Rotation = nil
	entity.Effects.Scale = nil
	entity.System = nil

	entity.Position.X = 0
	entity.Position.Y = 0
	entity.Rotation = 0.0
	entity.Scale = 1.0

	return &entity

}

func (entity *Entity) Update(now *time.Time) {

	if entity.Effects.Position != nil {

		result := entity.Effects.Position.Update(now)

		if result == false {
			entity.Effects.Position = nil
		}

	}

	if entity.Effects.Rotation != nil {

		result := entity.Effects.Rotation.Update(now)

		if result == false {
			entity.Effects.Rotation = nil
		}

	}

	if entity.Effects.Scale != nil {

		result := entity.Effects.Scale.Update(now)

		if result == false {
			entity.Effects.Scale = nil
		}

	}

}

func (entity *Entity) GetPosition() (int, int) {
	return entity.Position.X, entity.Position.Y
}

func (entity *Entity) GetRotation() float64 {
	return entity.Rotation
}

func (entity *Entity) GetScale() float64 {
	return entity.Scale
}

func (entity *Entity) SetPosition(x int, y int) {
	entity.Position.X = x
	entity.Position.Y = y
}

func (entity *Entity) SetRotation(rotation float64) {

	if rotation >= 0.0 && rotation <= 2 * math.Pi {
		entity.Rotation = rotation
	} else if rotation >= 0 {
		entity.Rotation = math.Mod(rotation, 2 * math.Pi)
	}

}

func (entity *Entity) SetScale(scale float64) {

	if scale >= 1.0 {
		entity.Scale = scale
	}

}

func (entity *Entity) TweenPosition(x int, y int) {

	effect := effects.Position{
		Entity:   entity,
		Start:    nil,
		Duration: 1000 * time.Millisecond,
		Tween:    effects.TweenBounceEaseOut,
		From:     [2]int{},
		To:       [2]int{x, y},
	}

	entity.Effects.Position = &effect

}

