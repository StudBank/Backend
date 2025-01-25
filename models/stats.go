package models

import "time"

type Stat struct {
	RecId string

	UserID int

	Subject int
	Value   string

	Created time.Time
	Start   time.Time
	End     time.Time
}
