package migrate

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/samsarahq/go/oops"
)

// Client is a simple wrapper around an http.Client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a new Client with a timeout of one minute.
func NewClient() *Client {
	return &Client{
		// BaseURL: "https://www.dnd5eapi.co/graphql",
		BaseURL: "http://mbp.local:3003/graphql",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) MakeGQLQuery(query string, v interface{}) error {
	req, err := c.buildHTTPRequest(query)
	if err != nil {
		return oops.Wrapf(err, "unable to build request")
	}

	return c.sendRequest(req, v)
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return oops.Wrapf(err, "unable to make request")
	}

	// TODO: better status code handling.
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return oops.Errorf("response error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return oops.Wrapf(err, "unable to decode response body")
	}

	return nil
}

func (c *Client) buildHTTPRequest(query string) (*http.Request, error) {

	req, err := http.NewRequest("POST", c.BaseURL, strings.NewReader(query))
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request")
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("apiKey", os.Getenv("REALM_API_KEY"))

	return req, nil
}
