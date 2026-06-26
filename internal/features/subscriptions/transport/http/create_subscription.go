package subscriptions_transport_http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateSubscriptionRequestDTO struct {
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	FinishDate  *time.Time `json:"finish_date"`
}

type CreateSubscriptionResponseDTO struct {
	ID          string     `json:"id"`
	Version     int        `json:"version"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	FinishDate  *time.Time `json:"finish_date"`
}

func (h *UsersHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateSubscriptionRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}
}
