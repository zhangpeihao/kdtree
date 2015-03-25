package kdtree

import (
	"errors"
)

var (
	ErrDimensionUnmatch = errors.New("Dimensions unmatch")
	ErrSearchStopped    = errors.New("Search stopped")
)
