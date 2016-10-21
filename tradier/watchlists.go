package tradier

import "encoding/json"

type WatchlistsService service

type Watchlists []*watchlist

type Watchlist struct {
	Name     *string `json:"name,omitempty"`
	ID       *string `json:"id,omitempty"`
	PublicID *string `json:"public_id,omitempty"`
	Items    Items   `json:"items,omitempty"`
}

type watchlist Watchlist

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
	var wlc struct {
		*watchlist `json:"watchlist,omitempty"`
	}

	if err := json.Unmarshal(b, &wlc); err == nil {
		*w = Watchlist(*wlc.watchlist)
		return nil
	}

	return nil
}

func (w *Watchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"watchlist": *w,
	})
}

func (w *Watchlists) UnmarshalJSON(b []byte) error {
	var wlc struct {
		W struct {
			W []*watchlist `json:"watchlist,omitempty"`
		} `json:"watchlists,omitempty"`
	}
	var wlObj struct {
		W struct {
			W *watchlist `json:"watchlist,omitempty"`
		} `json:"watchlists,omitempty"`
	}
	var wlNull string

	// log.Println(string(b))

	// If watchlist is null
	if err := json.Unmarshal(b, &wlNull); err == nil {
		return nil
	}

	// If watchlist is a JSON array
	if err := json.Unmarshal(b, &wlc); err == nil {
		*w = Watchlists(wlc.W.W)
		// *w = wlc.Watchlists
		return nil
	}

	// If watchlist is a single object
	if err := json.Unmarshal(b, &wlObj); err == nil {
		wl := make([]*watchlist, 0)
		wl = append(wl, wlObj.W.W)
		*w = Watchlists(wl)
		return nil
	}

	return nil
}
