package lcp

import "time"

type Response[T any] struct {
	Data    T
	Updated time.Time
}
