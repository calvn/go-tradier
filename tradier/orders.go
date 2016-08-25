package tradier

import (
	"encoding/json"
	"time"
)

// `orders` JSON can be an object or string literal "null"
type Orders struct {
	Order  []Order `json:"order,omitempty"`
	isNull *string
}

type orders Orders

// `order` JSON can be an object or array of objects
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
	// orderArray        `json:"-"` // Used internally to store array of `order` object as slice
}

type order struct {
	*Order `json:"order,omitempty"`
}

type orderArray []order

func (o *Orders) UnmarshalJSON(b []byte) (err error) {
	ordersStr := ""
	ordersObj := orders{}
	orderObj := order{}

	// If order is a string, i.e. "null"
	if err = json.Unmarshal(b, &ordersStr); err == nil {
		o.isNull = &ordersStr
		return nil
	}

	// If order is an array
	if err = json.Unmarshal(b, &ordersObj); err == nil {
		*o = Orders(ordersObj)
		return nil
	}

	// If order is an object
	if err = json.Unmarshal(b, &orderObj); err == nil {
		tmp := []Order{*orderObj.Order}
		*o = Orders{Order: tmp}
		return nil
	}

	return nil
}

func (o *Orders) MarshalJSON() ([]byte, error) {
	// If order is null
	if o.isNull != nil {
		return json.Marshal(o.isNull)
	}

	// If order array is size 1, return object directly
	if len(o.Order) == 1 {
		return json.Marshal(map[string]interface{}{
			"order": o.Order[0],
		})
	}

	return json.Marshal(*o)
}

// func (o *Order) UnmarshalJSON(b []byte) (err error) {
// 	orderObj := order{}
//
// 	log.Println(string(b))
//
// 	// If order is an object
// 	if err = json.Unmarshal(b, &orderObj); err == nil {
// 		log.Println("here")
// 		*o = Order(orderObj)
// 		return nil
// 	}
// 	//
// 	// // If order is an array
// 	// if err = json.Unmarshal(b, &orderArr); err == nil {
// 	// 	o.orderArray = orderArr
// 	// 	return nil
// 	// } else {
// 	// 	log.Println(err) // FIXME: Better error handling
// 	// }
//
// 	return nil
// }

// func (o *Order) MarshalJSON() ([]byte, error) {
//
// 	if o.orderArray != nil {
// 		return json.Marshal(o.orderArray)
// 	}
//
// 	return json.Marshal(Order(*o))
// }
