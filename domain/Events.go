package domain

import "time"

type Events struct {
	Name        string
	Time        time.Duration
	Description string
}
