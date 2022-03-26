/*
Copyright © 2022 Kyle Chadha @kylechadha
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kylechadha/recreation-gov-notify/notify"

	"github.com/inconshreveable/log15"
)

func runNotify(cfg *notify.Config) {
	log := log15.New()
	if !cfg.Debug {
		log.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StdoutHandler))
	}

	app := notify.New(log, cfg)
	reader := bufio.NewReader(os.Stdin)

	var campground notify.Campground
Outer:
	for {
		fmt.Println("Which campground are you looking for?")

		reader.Reset(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Sorry, there was an error, please try again. Error : %w\n", err)
			continue
		}
		query := strings.Replace(input, "\n", "", -1) // convert CRLF to LF

		campgrounds, err := app.Search(query)
		if err != nil {
			fmt.Printf("Sorry, there was an error, please try again. Error : %w\n", err)
			continue
		}
		if len(campgrounds) == 0 {
			fmt.Println("Sorry, we didn't find any campgrounds for that query. Please try again")
			continue
		}

		fmt.Println("Select the number that best matches:")
		for i, c := range campgrounds {
			fmt.Printf("[%d] %s\n", i+1, c.Name)
		}
		lastIndex := len(campgrounds) + 1
		fmt.Printf("[%d] None of these, let me search again\n", lastIndex)

		for {
			reader.Reset(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Sorry, there was an error, please try again. Error : %w\n", err)
				continue
			}
			choice, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))
			if err != nil || choice > lastIndex {
				fmt.Printf("Sorry, that was an invalid selection, please try again")
				continue
			}
			if choice == lastIndex {
				continue Outer
			}

			campground = campgrounds[choice-1]
			break Outer
		}
	}

	var checkInDate string
	var start time.Time
	for {
		fmt.Println(`When's your check in? Please enter in "MM-DD-YYYY" format.`)

		reader.Reset(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Sorry, there was an error, please try again. Error : %w\n", err)
			continue
		}
		checkInDate = strings.Replace(input, "\n", "", -1) // convert CRLF to LF

		start, err = time.Parse("01-02-2006", checkInDate)
		if err != nil {
			fmt.Println("Sorry I couldn't parse that date. please try again. Error : %w\n", err)
			continue
		}
		break
	}

	var checkOutDate string
	var end time.Time
	for {
		fmt.Println(`When's your check out? Please enter in "MM-DD-YYYY" format.`)

		reader.Reset(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Sorry, there was an error, please try again. Error : %w\n", err)
			continue
		}
		checkOutDate = strings.Replace(input, "\n", "", -1) // convert CRLF to LF

		endRaw, err := time.Parse("01-02-2006", checkOutDate)
		if err != nil {
			fmt.Println("Sorry I couldn't parse that date. please try again. Error : %w\n", err)
			continue
		}
		end = endRaw.AddDate(0, 0, -1) // checkOutDate does not need to be available

		if start.After(end) {
			fmt.Println("Check out needs to be after check in ;)")
			continue
		}
	}

	fmt.Printf("Alright! Searching recreation.gov availability for %s from %s to %s\n", campground.Name, checkInDate, checkOutDate)
}
