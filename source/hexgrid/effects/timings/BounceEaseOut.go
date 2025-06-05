package timings

func BounceEaseOut(t float64) float64 {

	if t < (1.0 / 2.75) {

		return 7.5625 * t * t

	} else if t < (2.0 / 2.75) {

		k := t - (1.5 / 2.75)

		return 7.5625 * k * k + 0.75

	} else if t < (2.5 / 2.75) {

		k := t - (2.25 / 2.75)

		return 7.5625 * k * k + 0.9375

	} else {

		k := t - (2.625 / 2.75)

		return 7.5625 * k * k + 0.984375

	}

}
