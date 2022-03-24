package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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

// ** HERE
// then connect to text/email
// then convert to CLI app
// - Split out front end vs back end here
// - How do we programmatically determine campground ID? Is there another API we can hit to search?
// then create web frontend
// productionize w/ cloud functions?
// - Do this when you built the front end, either we have one backend app that spins up goroutines to poll
// Or spins up a cloud function

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	l := log15.New()
	client := http.Client{}

	debug := true
	if !debug {
		l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StdoutHandler))
	}

	campground := 231969
	checkInDate := "2022-06-01"
	checkOutDate := "2022-06-02"

	l.Info("Searching recreation.gov...", "campground", campground, "checkIn", checkInDate, "checkOut", checkOutDate)

	st, err := time.Parse("2006-01-02", checkInDate)
	if err != nil {
		l.Error("Invalid check in date", "err", err)
		exitCode = 1
		return
	}
	etRaw, err := time.Parse("2006-01-02", checkOutDate)
	if err != nil {
		l.Error("Invalid check out date", "err", err)
		exitCode = 1
		return
	}
	et := etRaw.AddDate(0, 0, -1) // checkOutDate does not need to be available

	if st.After(et) {
		l.Error("Start date is after end date")
		exitCode = 1
		return
	}

	curPeriod := fmt.Sprintf("%d-%02d", st.Year(), st.Month())
	endPeriod := fmt.Sprintf("%d-%02d", et.Year(), et.Month())

	var months []string
	months = append(months, curPeriod)

	initial := st
	for curPeriod != endPeriod {
		st = st.AddDate(0, 1, 0)
		curPeriod = fmt.Sprintf("%d-%02d", st.Year(), st.Month())
		months = append(months, curPeriod)
	}
	st = initial

	available := make(map[string]map[string]bool)
	baseURL := fmt.Sprintf("https://www.recreation.gov/api/camps/availability/campground/%d/month", campground)
	for _, m := range months {
		l.Debug("Requesting data", "month", m)

		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		q := url.Values{}
		q.Add("start_date", m+"-01T00:00:00.000Z")
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
					if available[c.Site] == nil {
						available[c.Site] = make(map[string]bool)
					}

					available[c.Site][date] = true
				}
			}
		}
	}

	var results []string
Outer:
	for site, dates := range available {
		st = initial
		for !st.After(et) {
			date := fmt.Sprintf("%sT00:00:00Z", st.Format("2006-01-02"))
			if dates[date] {
				l.Debug(fmt.Sprintf("Site %s available for %s", site, st.Format("2006-01-02")))
				st = st.AddDate(0, 0, 1)
			} else {
				continue Outer
			}
		}

		l.Info(fmt.Sprintf("Site %s is available", site))
		results = append(results, site)
	}

	if len(results) == 0 {
		l.Info("Sorry, no available campsites were found for the full date range")
	}
}
