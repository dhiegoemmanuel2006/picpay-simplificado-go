package request

import (
	"github.com/google/uuid"
)

type TranferRequest struct {
	Value float64
	Payer uuid.UUID
	Payee uuid.UUID
}
