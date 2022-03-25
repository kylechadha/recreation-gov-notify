/*
Copyright © 2022 Kyle Chadha @kylechadha
*/
package cmd

import (
	"github.com/kylechadha/recreation-gov-notify/notify"

	"github.com/inconshreveable/log15"
)

func runNotify(cfg *notify.Config) {
	l := log15.New()
	if !cfg.Debug {
		l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StdoutHandler))
	}
}
