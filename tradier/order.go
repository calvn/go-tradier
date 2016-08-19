package tradier

import (
	"encoding/json"
	"time"
)

// `orders` can be an object or string literal "null"
type Orders struct {
	Order *Order `json:"order,omitempty"`
}

// `order` can be an object or array of objects
type Order struct {
	AvgFillPrice      *float64   `json:"avg_fill_price,omitempty"`
	Class             *string    `json:"class,omitempty"`
	CreateDate        *time.Time `json:"create_date,omitempty"`
	Duration          *string    `json:"duration,omitempty"`
	ExecQuantity      *int       `json:"exec_quantity,omitempty"`
	ID                *int       `json:"id,omitempty"`
	LastFillPrice     *float64   `json:"last_fill_price,omitempty"`
	LastFillQuantity  *int       `json:"last_fill_quantity,omitempty"`
	Quantity          *float64   `json:"quantity,omitempty"`
	RemainingQuantity *int       `json:"remaining_quantity,omitempty"`
	Side              *string    `json:"side,omitempty"`
	Status            *string    `json:"status,omitempty"`
	Symbol            *string    `json:"symbol,omitempty"`
	TransactionDate   *time.Time `json:"transaction_date,omitempty"`
	Type              *string    `json:"type,omitempty"`
}

type order Order

type orders Orders

func (o *Orders) UnmarshalJSON(b []byte) (err error) {
	ordersStr := ""
	ordersObj := orders{}

	if err = json.Unmarshal(b, &ordersObj); err == nil {
		*o = Orders(ordersObj)
		return
	}
	if err = json.Unmarshal(b, &ordersStr); err == nil {
		o.Order = &Order{}
		o.Order.Type = &ordersStr
		return
	}

	return nil
}

func (o *Orders) MarshalJSON() ([]byte, error) {
	if o.Order != nil && *o.Order.Type == "null" {
		return json.Marshal(o.Order.Type)
	}

	return json.Marshal(Orders(*o))
}

// func (o *Order) UnmarshalJSON(b []byte) (err error) {
// 	orderStr := ""
// 	// var orderObj *order
// 	orderArr := []interface{}{
// 		&o.AvgFillPrice,
// 		&o.Class,
// 		&o.CreateDate,
// 		&o.Duration,
// 		&o.ExecQuantity,
// 		&o.ID,
// 		&o.LastFillPrice,
// 		&o.LastFillQuantity,
// 		&o.Quantity,
// 		&o.RemainingQuantity,
// 		&o.Side,
// 		&o.Status,
// 		&o.Symbol,
// 		&o.TransactionDate,
// 		&o.Type,
// 	}
//
// 	// If order is a string
// 	if err = json.Unmarshal(b, &orderStr); err == nil {
// 		o.Type = &orderStr
// 		return
// 	}
//
// 	// If order is an array
// 	if err = json.Unmarshal(b, orderArr); err == nil {
// 		return nil
// 	}
//
// 	return nil
// }
