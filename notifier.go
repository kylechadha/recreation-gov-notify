package main

import (
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Notifier interface {
	Notify(to string, available []string) error
}

type SMSNotifier struct {
	l      log15.Logger
	client *twilio.RestClient
	from   string
}

func NewSMSNotifier(l log15.Logger, accountSid string, authToken string, from string) *SMSNotifier {
	return &SMSNotifier{
		l: l,
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
		from: from,
	}
}

func (n SMSNotifier) Notify(to string, available []string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(n.from)
	params.SetBody(fmt.Sprintf("Good news from the recreation.gov notifier! The following sites are available: %s", strings.Join(available, ", ")))

	resp, err := n.client.ApiV2010.CreateMessage(params)
	if err != nil {
		return err
	}

	n.l.Debug("SMS message sent", "status", resp.Status, "to", to)
	return nil
}
