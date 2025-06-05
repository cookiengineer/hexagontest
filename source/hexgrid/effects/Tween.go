package effects

type Tween int

const (
	TweenLinear Tween = iota
	TweenEaseIn
	TweenEaseOut
	TweenBounceEaseIn
	TweenBounceEaseOut
)
