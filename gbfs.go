package gbfs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// GBFS represent a single system or geographic area
// in which vehicles are operated.
type GBFS struct {
	Header
	Data LanguageFeeds `json:"data"`
}

// Every JSON file presented in the specification
// contains the same common header information
// at the top level of the JSON response object.
type Header struct {
	LastUpdated Timestamp `json:"last_updated"`
	TTL         int       `json:"ttl"`
	Version     string    `json:"version"` // added in v1.1
}

type LanguageFeeds map[string]DataFeeds

type DataFeeds struct {
	Feeds []Feed `json:"feeds"`
}

type Feed struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) LoadGBFS(url string) (*GBFS, error) {
	resp, err := c.sendRequest(url)
	if err != nil {
		return nil, errors.Wrap(err, "send request")
	}
	defer resp.Body.Close()

	var gbfs GBFS

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if err := json.Unmarshal(b, &gbfs); err != nil {
		return nil, errors.Wrap(err, "unmarshal JSON")
	}

	return &gbfs, nil
}
