package tradier

import "fmt"

func (w *WatchlistsService) Get(watchlistId string) (*Watchlist, *Response, error) {
	u := fmt.Sprintf("watchlists/%s", watchlistId)

	req, err := w.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	wl := &Watchlist{}

	resp, err := w.client.Do(req, wl)
	if err != nil {
		return nil, resp, err
	}

	return wl, resp, nil
}
