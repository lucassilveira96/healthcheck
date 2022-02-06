package models

import "time"

type Docker struct {
	Name      string
	Running   bool
	CheckedAt time.Time
}
