package tradier

func (w *WatchlistsService) All() (*Watchlists, *Response, error) {
	req, err := w.client.NewRequest("GET", "watchlists", nil)
	if err != nil {
		return nil, nil, err
	}

	wl := &Watchlists{}

	resp, err := w.client.Do(req, wl)
	if err != nil {
		return nil, resp, err
	}

	return wl, resp, nil
}
