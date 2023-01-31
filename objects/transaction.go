package objects

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	id        uuid.UUID `json:"id"`
	amount    int       `json:"amount"`
	from      uuid.UUID `json:"from"`
	to        uuid.UUID `json:"to"`
	createdAt time.Time `json:"created_at"`
}
