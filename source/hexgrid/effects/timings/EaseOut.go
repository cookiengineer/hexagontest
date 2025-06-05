package timings

func EaseOut(t float64) float64 {
	return t * (2.0 - t)
}
