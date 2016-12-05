package tradier

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const (
	libraryVersion = "0.0.1"
	defaultBaseURL = "https://api.tradier.com/v1/"
	userAgent      = "go-tradier/" + libraryVersion
)

// Client takes care of managing communication to the Tradier api
type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	BaseURL *url.URL

	UserAgent string

	common service

	User       *UserService
	Account    *AccountService
	Order      *OrderService
	Watchlists *WatchlistsService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.common.client = c
	c.User = (*UserService)(&c.common)
	c.Account = (*AccountService)(&c.common) // FIXME: should be AccountsService
	c.Order = (*OrderService)(&c.common)     // FIXME: should be OrdersService
	c.Watchlists = (*WatchlistsService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	// If the body is not empty, assume it's a form data
	var buf io.ReadWriter
	if body != nil {
		buf = bytes.NewBufferString(body.(string))
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Always request JSON
	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

type Response struct {
	*http.Response
	// Rate
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	// response.Rate = parseRate(r)
	return response
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	// TODO: Do rate limit checking

	// b, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// log.Printf("%s", b)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		var e error
		b, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			return nil, e
		}
		log.Println(string(b))
		return nil, err
	}

	// log.Println(req.FormValue("symbol"))
	//
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// log.Printf("%s", b)

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func Float64(v float64) *float64 { return &v }
