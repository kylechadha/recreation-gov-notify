package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/inconshreveable/log15"
)

type Campsites struct {
	Campsites map[string]Campsite `json:"campsites"`
	Count     int                 `json:"count"`
}

type Campsite struct {
	Availabilities      map[string]string `json:"availabilities"`
	CampsiteID          string            `json:"campsite_id"`
	CampsiteReserveType string            `json:"campsite_reserve_type"`
	CampsiteType        string            `json:"campsite_type"`
	Loop                string            `json:"loop"`
	MaxNumPeople        int               `json:"max_num_people"`
	MinNumPeople        int               `json:"min_num_people"`
	Site                string            `json:"site"`
	TypeOfUse           string            `json:"type_of_use"`
}

func main() {
	l := log15.New()
	client := http.Client{}

	// ** HERE
	// then visual display out
	// then connect to text/email
	// then convert to CLI app

	debug := true
	campground := 231959
	startDate := "2021-12-27"
	endDate := "2022-01-03"

	l.Info(fmt.Sprintf("Searching recreation.gov for availability from %s to %s", startDate, endDate))

	if !debug {
		l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StdoutHandler))
	}

	st, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Fatal(err)
	}
	et, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Fatal(err)
	}

	startPeriod := fmt.Sprintf("%d-%02d-01", st.Year(), st.Month())
	endPeriod := fmt.Sprintf("%d-%02d-01", et.Year(), et.Month())

	var months []string
	months = append(months, startPeriod)

	if startPeriod != endPeriod {
		for st.Before(et) {
			st = st.AddDate(0, 1, 0)
			months = append(months, fmt.Sprintf("%d-%02d-01", st.Year(), st.Month()))
		}
	}

	available := make(map[string]map[string]bool)
	baseURL := fmt.Sprintf("https://www.recreation.gov/api/camps/availability/campground/%d/month", campground)
	for _, m := range months {
		l.Debug(fmt.Sprintf("Requesting data for %s", strings.TrimSuffix(m, "-01")))

		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		q := url.Values{}
		q.Add("start_date", m+"T00:00:00.000Z")
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

		for _, c := range cs.Campsites {
			for date, a := range c.Availabilities {
				if a == "Available" {

				}
			}
		}
	}
}
