package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Campsites struct {
	Campsites map[string]Campsite `json:"campsites"`
	Count     int                 `json:"count"`
}

type Campsite struct {
	Availabilities      map[string]Availability `json:"availabilities"`
	CampsiteID          string                  `json:"campsite_id"`
	CampsiteReserveType string                  `json:"campsite_reserve_type"`
	CampsiteType        string                  `json:"campsite_type"`
	Loop                string                  `json:"loop"`
	MaxNumPeople        int                     `json:"max_num_people"`
	MinNumPeople        int                     `json:"min_num_people"`
	Site                string                  `json:"site"`
	TypeOfUse           string                  `json:"type_of_use"`
}

type Availability string

const (
	Unknown       Availability = "Unknown"
	Available     Availability = "Available"
	NotAvailable  Availability = "Not Available"
	Reserved      Availability = "Reserved"
	NotReservable Availability = "Not Reservable"
)

func main() {
	client := http.Client{}

	// ** HERE
	// start building params in part
	// then visual display out
	// then connect to text/email
	// then make CLI app

	campground := 231959
	baseURL := fmt.Sprintf("https://www.recreation.gov/api/camps/availability/campground/%d/month", campground)
	startDate := "2021-12-01"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := url.Values{}
	q.Add("start_date", startDate+"T00:00:00.000Z")
	req.URL.RawQuery = q.Encode()

	// Need to spoof the user agent or CloudFront blocks us.
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(resp.StatusCode, ": ", string(bytes))
	}

	var cs Campsites
	err = json.NewDecoder(resp.Body).Decode(&cs)
	if err != nil {
		log.Fatal(err)
	}

	var dates []string
	for _, c := range cs.Campsites {
		for date, a := range c.Availabilities {
			if a == Available {
				dates = append(dates, strings.TrimSuffix(date, "T00:00:00Z"))
			}
		}
	}
	fmt.Println(dates)
}
