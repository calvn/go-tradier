package tradier

import (
	"encoding/json"
	"log"
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
	ExecQuantity      *float64   `json:"exec_quantity,omitempty"`
	ID                *int       `json:"id,omitempty"`
	LastFillPrice     *float64   `json:"last_fill_price,omitempty"`
	LastFillQuantity  *float64   `json:"last_fill_quantity,omitempty"`
	Quantity          *float64   `json:"quantity,omitempty"`
	RemainingQuantity *float64   `json:"remaining_quantity,omitempty"`
	Side              *string    `json:"side,omitempty"`
	Status            *string    `json:"status,omitempty"`
	Symbol            *string    `json:"symbol,omitempty"`
	TransactionDate   *time.Time `json:"transaction_date,omitempty"`
	Type              *string    `json:"type,omitempty"`
	orderArray        []order    // Used internally to store array of `order` object
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
	if o.Order != nil && o.Order.orderArray == nil && *o.Order.Type == "null" {
		return json.Marshal(o.Order.Type)
	}

	return json.Marshal(Orders(*o))
}

func (o *Order) UnmarshalJSON(b []byte) (err error) {
	orderObj := order{}
	orderArr := []order{}

	// log.Println(string(b))

	// If order is an object
	if err = json.Unmarshal(b, &orderObj); err == nil {
		*o = Order(orderObj)
		return
	}

	// If order is an array
	if err = json.Unmarshal(b, &orderArr); err == nil {
		o.orderArray = orderArr
		return nil
	} else {
		log.Println(err)
	}

	return nil
}

func (o *Order) MarshalJSON() ([]byte, error) {

	if o.orderArray != nil {
		return json.Marshal(o.orderArray)
	}

	return json.Marshal(Order(*o))
}
