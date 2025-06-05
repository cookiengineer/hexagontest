package effects

import "battlemap/hexgrid/effects/timings"
import "time"

type RotationEffectee interface {
	GetRotation() float64
	SetRotation(float64)
}

type Rotation struct {
	Entity   RotationEffectee `json:"entity"`
	Start    *time.Time       `json:"start"`
	End      *time.Time       `json:"start"`
	Duration time.Duration    `json:"duration"`
	Tween    Tween            `json:"tween"`
	From     float64          `json:"from"`
	To       float64          `json:"to"`
}

func (effect *Rotation) Update(now *time.Time) bool {

	var result bool

	if effect.Entity != nil {

		if effect.Start == nil {

			from := effect.Entity.GetRotation()
			end := now.Add(effect.Duration)

			effect.Start = now
			effect.End = &end
			effect.From = from

			effect.Entity.SetRotation(effect.From)
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

			current := effect.From + float64((effect.To - effect.From) * t)

			effect.Entity.SetRotation(current)
			result = true

		} else if now.After(*effect.End) {

			effect.Entity.SetRotation(effect.To)

		}

	}

	return result

}
