package tradier

import (
	"encoding/json"
	"log"
)

type WatchlistsService service

type Watchlists []*Watchlist

// Private struct used to unmarshal JSON and hold the object a struct
type watchlists struct {
	Collection *struct {
		Watchlist []Watchlist `json:"watchlist,omitempty"`
	} `json:"watchlists,omitempty"`
}

type Watchlist struct {
	Name     *string `json:"name,omitempty"`
	ID       *string `json:"id,omitempty"`
	PublicID *string `json:"public_id,omitempty"`
	Items    Items   `json:"items,omitempty"`
}

type watchlist Watchlist

type watchlistContainer struct {
	*watchlist `json:"watchlist,omitempty"`
}

type Items struct {
	Item []*WatchlistItem `json"item,omitempty"`
}

type items Items

type WatchlistItem struct {
	Symbol *string `json:"symbol,omitempty"`
	ID     *string `json:"id,omitempty"`
}

type watchlistItem struct {
	*WatchlistItem `json:"item,omitempty"`
}

func (i *Items) UnmarshalJSON(b []byte) error {
	itemsStr := ""
	itemsObj := items{}
	itemObj := watchlistItem{}

	// If items is a string, i.e. "null"
	if err := json.Unmarshal(b, &itemsStr); err == nil {
		return nil
	}

	// If itemsr is a JSON array
	if err := json.Unmarshal(b, &itemsObj); err == nil {
		*i = Items(itemsObj)
		return nil
	}

	// If items is an object
	if err := json.Unmarshal(b, &itemObj); err == nil {
		obj := WatchlistItem(*itemObj.WatchlistItem)
		var slice []*WatchlistItem
		slice = append(slice, &obj)
		*i = Items{Item: slice}
		return nil
	}

	return nil
}

func (i *Items) MarshalJSON() ([]byte, error) {
	if len(i.Item) == 0 {
		return json.Marshal("null")
	}

	if len(i.Item) == 1 {
		return json.Marshal(map[string]interface{}{
			"items": i.Item[0],
		})
	}

	return json.Marshal(*i)
}

// Unmarshal json into Watchlist object
func (w *Watchlist) UnmarshalJSON(b []byte) error {
	wlc := watchlistContainer{}

	if err := json.Unmarshal(b, &wlc); err == nil {
		*w = Watchlist(*wlc.watchlist)
		return nil
	} else {
		log.Println(err)
		return err
	}

	return nil
}

func (w *Watchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"watchlist": *w,
	})
}
