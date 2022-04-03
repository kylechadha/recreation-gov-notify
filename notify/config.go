package notify

import "time"

type Config struct {
	Debug        bool
	PollInterval time.Duration
	SMSFrom      string
	EmailFrom    string
}
