package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int
	Version     int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	DateStart   time.Time
	DateFinish  *time.Time
}
