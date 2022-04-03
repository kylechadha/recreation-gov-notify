/*
Copyright © 2022 Kyle Chadha @kylechadha
*/
package notify

import (
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSNotifier struct {
	log        log15.Logger
	client     *twilio.RestClient
	fromNumber string
}

func NewSMSNotifier(log log15.Logger, fromNumber string) *SMSNotifier {
	return &SMSNotifier{
		log:        log,
		fromNumber: fromNumber,
		client:     twilio.NewRestClient(),
	}
}

const SMSTemplate = `
Good news from the (very unofficial) Recreation.gov Notifier!

The following sites are available for '%s' from %s to %s:

%s`

func (n *SMSNotifier) Notify(to string, campgroundName, checkInDate, checkOutDate string, available []string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(n.fromNumber)

	sites := " - Site " + strings.Join(available, "\n - Site ")
	params.SetBody(fmt.Sprintf(SMSTemplate, campgroundName, checkInDate, checkOutDate, sites))

	resp, err := n.client.ApiV2010.CreateMessage(params)
	if err != nil {
		return err
	}

	n.log.Debug("SMS message sent", "status", *resp.Status, "to", to)
	return nil
}
