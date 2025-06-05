package hexgrid

import "time"

type Effect interface {
	Update(*time.Time) bool
}
