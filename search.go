package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const suggestURL = "https://www.recreation.gov/api/search/suggest?q=%s&geocoder=true"

type SuggestResponse struct {
	InventorySuggestions []Campground `json:"inventory_suggestions"`
}

func Suggest(query string) ([]Campground, error) {
	client := http.Client{} // ** Create a recreationGovClient type

	url := fmt.Sprintf(suggestURL, url.QueryEscape(query))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Need to spoof the user agent or CloudFront blocks us.
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s: %s", resp.Status, string(bytes))
	}

	var sr SuggestResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return nil, err
	}

	return sr.InventorySuggestions, nil
}
