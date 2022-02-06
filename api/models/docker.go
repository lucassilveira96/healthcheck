package models

import "time"

type Docker struct {
	Name      string
	Status    bool
	CheckedAt time.Time
}
