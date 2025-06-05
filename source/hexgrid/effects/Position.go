package effects

import "battlemap/hexgrid/effects/timings"
import "time"

type PositionEffectee interface {
	GetPosition() (int, int)
	SetPosition(int, int)
}

type Position struct {
	Entity   PositionEffectee `json:"entity"`
	Start    *time.Time       `json:"start"`
	End      *time.Time       `json:"start"`
	Duration time.Duration    `json:"duration"`
	Tween    Tween            `json:"tween"`
	From     [2]int           `json:"from"`
	To       [2]int           `json:"to"`
}

func (effect *Position) Update(now *time.Time) bool {

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

			var t float64 = float64(now.Sub(*effect.Start)) / float64(effect.Duration)

			if effect.Tween == TweenLinear {
				t = timings.Linear(t)
			} else if effect.Tween == TweenEaseIn {
				t = timings.EaseIn(t)
			} else if effect.Tween == TweenEaseOut {
				t = timings.EaseOut(t)
			} else if effect.Tween == TweenBounceEaseIn {
				t = timings.BounceEaseIn(t)
			} else if effect.Tween == TweenBounceEaseOut {
				t = timings.BounceEaseOut(t)
			}

			current_x := effect.From[0] + int((float64(effect.To[0]) - float64(effect.From[0])) * t)
			current_y := effect.From[1] + int((float64(effect.To[1]) - float64(effect.From[1])) * t)

			effect.Entity.SetPosition(current_x, current_y)
			result = true

		} else if now.After(*effect.End) {

			effect.Entity.SetPosition(effect.To[0], effect.To[1])

		}

	}

	return result

}
