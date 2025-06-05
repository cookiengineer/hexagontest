package hexgrid

import "time"

type PositionEffectee interface {
	GetPosition() (int, int)
	SetPosition(int, int)
}

type PositionEffect struct {
	Entity   PositionEffectee `json:"entity"`
	Start    *time.Time       `json:"start"`
	End      *time.Time       `json:"start"`
	Duration time.Duration    `json:"duration"`
	From     [2]int           `json:"from"`
	To       [2]int           `json:"to"`
}

func (effect *PositionEffect) Update(now *time.Time) bool {

	var result bool

	if effect.Entity != nil {

		if effect.Start == nil {

			from_x, from_y := effect.Entity.GetPosition()
			end := now.Add(effect.Duration)

			effect.Start = now
			effect.End = &end
			effect.From = [2]int{from_x, from_y}
			effect.Entity.SetPosition(effect.From[0], effect.From[1])

			result = true

		} else if now.After(*effect.Start) && now.Before(*effect.End) {

			t := float64(now.Sub(*effect.Start)) / float64(effect.Duration)
			current_x := effect.From[0] + int((float64(effect.To[0]) - float64(effect.From[0])) * t)
			current_y := effect.From[1] + int((float64(effect.To[1]) - float64(effect.From[1])) * t)

			effect.Entity.SetPosition(current_x, current_y)

		} else if now.After(*effect.End) {

			effect.Entity.SetPosition(effect.To[0], effect.To[1])

			result = false

		}

	}

	return result

}
