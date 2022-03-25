package notify

import "time"

type Config struct {
	Debug        bool
	PollInterval time.Duration

	CampgroundID int
	CheckInDate  string
	CheckOutDate string
}
