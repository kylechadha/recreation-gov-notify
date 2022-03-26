package notify

import "time"

type Config struct {
	Debug        bool
	PollInterval time.Duration

	CampgroundID string
	CheckInDate  string
	CheckOutDate string
}
