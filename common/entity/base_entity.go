package entity

import "time"

type BaseEntity struct {
	CreatedAt time.Time
	UpdateAt  time.Time
}
