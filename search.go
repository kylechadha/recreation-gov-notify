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

type Campground struct {
	CampsiteReserveType []string `json:"campsite_reserve_type"`
	CampsiteTypeOfUse   []string `json:"campsite_type_of_use"`
	City                string   `json:"city"`
	CountryCode         string   `json:"country_code"`
	EntityID            string   `json:"entity_id"`
	EntityType          string   `json:"entity_type"`
	IsInventory         bool     `json:"is_inventory"`
	Lat                 string   `json:"lat"`
	Lng                 string   `json:"lng"`
	Name                string   `json:"name"`
	OrgID               string   `json:"org_id"`
	OrgName             string   `json:"org_name"`
	ParentEntityID      string   `json:"parent_entity_id"`
	ParentEntityType    string   `json:"parent_entity_type"`
	ParentName          string   `json:"parent_name"`
	PreviewImageURL     string   `json:"preview_image_url"`
	Reservable          bool     `json:"reservable"`
	StateCode           string   `json:"state_code"`
	Text                string   `json:"text"`
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
