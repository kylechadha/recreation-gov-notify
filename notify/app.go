/*
Copyright © 2022 Kyle Chadha @kylechadha
*/
package notify

import "github.com/inconshreveable/log15"

type App struct {
	cfg *Config
	log log15.Logger

	client    *Client
	notifiers []Notifier
}

type Notifier interface {
	// Notify(to string, campground, checkInDate, checkOutDate string, available []string) error
	Notify(to string, available []string) error
}

func New(log log15.Logger, cfg *Config) *App {
	return &App{
		cfg:    cfg,
		log:    log,
		client: NewClient(log),
	}
}

func (a *App) Search(query string) ([]Campground, error) {
	return a.client.Search(query)
}

func (a *App) Poll() {

}

func (a *App) Notify() {

}
