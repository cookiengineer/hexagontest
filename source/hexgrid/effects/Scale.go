package effects

import "battlemap/hexgrid/effects/timings"
import "time"

type ScaleEffectee interface {
	GetScale() float64
	SetScale(float64)
}

type Scale struct {
	Entity   ScaleEffectee `json:"entity"`
	Start    *time.Time    `json:"start"`
	End      *time.Time    `json:"start"`
	Duration time.Duration `json:"duration"`
	Tween    Tween         `json:"tween"`
	From     float64       `json:"from"`
	To       float64       `json:"to"`
}

func (effect *Scale) Update(now *time.Time) bool {

	var result bool

	if effect.Entity != nil {

		if effect.Start == nil {

			from := effect.Entity.GetScale()
			end := now.Add(effect.Duration)

			effect.Start = now
			effect.End = &end
			effect.From = from

			effect.Entity.SetScale(effect.From)
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

			effect.Entity.SetScale(current)
			result = true

		} else if now.After(*effect.End) {

			effect.Entity.SetScale(effect.To)

		}

	}

	return result

}
