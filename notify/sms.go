package notify

import (
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSNotifier struct {
	l      log15.Logger
	client *twilio.RestClient
	from   string
}

func NewSMSNotifier(l log15.Logger, from string) *SMSNotifier {
	return &SMSNotifier{
		l:      l,
		client: twilio.NewRestClient(),
		from:   from,
	}
}

const SMSTemplate = `
Good news from the (very unofficial) Recreation.gov Notifier!

The following sites are available for '%s' from %s to %s:
%s`

func (n SMSNotifier) Notify(to string, campground, checkInDate, checkOutDate string, available []string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(n.from)

	sites := "\n - Site " + strings.Join(available, "\n - Site ")
	params.SetBody(fmt.Sprintf(SMSTemplate, campground, checkInDate, checkOutDate, sites))

	resp, err := n.client.ApiV2010.CreateMessage(params)
	if err != nil {
		return err
	}

	n.l.Debug("SMS message sent", "status", *resp.Status, "to", to)
	return nil
}
