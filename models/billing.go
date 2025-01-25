package models

import "time"

const (
	SUB_STSTUS_ACTIVE = iota + 1
	SUB_STSTUS_CANCELED_ACTIVE
	SUB_STSTUS_CANCELED
	SUB_STSTUS_PAUSED_BLOCKED_CARD
	SUB_STSTUS_PAUSED
	SUB_STSTUS_NEVER_BOUGHT
)

type Subscription struct {
	UserID int

	Status int

	LastWithdraw time.Time
	KassaCode    string
	Card         string

	Untill time.Time

	UsageK       int
	TokensWasted int64
}

type Withdawl struct {
	ID     string
	UserID int

	Amount      int
	KassaStatus int

	TimeInitiated time.Time
	TimeConfirmed time.Time
	TimeChanged   time.Time
	TimeFinalized time.Time
}
