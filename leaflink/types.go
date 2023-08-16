package leaflink

import (
	"time"
)

type Order struct {
	Number    string    `json:"number"`
	CreatedOn time.Time `json:"created_on"`
}

type OrderSalesRep struct {
	ID        int       `json:"id"`
	Order     string    `json:"order"`
	Rep       int       `json:"rep"`
	CreatedOn time.Time `json:"created_on"`
}
