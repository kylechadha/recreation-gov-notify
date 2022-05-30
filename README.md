# Recreation.gov Campsite Availability

Get notified when a campsite you're interested in becomes available on [recreation.gov](https://recreation.gov). Campsites book up early, but cancelations are common and recreation.gov has no built-in notification feature.


## Build

```shell
go build
```

## Configure

edit the yaml file `~/.rgn`
```yaml
Debug: false

# poll interval in nanos
PollInterval: 30000000000

# send SMS notification to that number 
SMSTo: "+11234567890"

# send email notification to that number 
EmailTo: "someone@acme.com"
```

### SMS
In addition to the `SMSTo` configuration, setup a https://www.twilio.com/ account

Setup env variables
`TWILIO_AUTH_TOKEN`
`TWILIO_ACCOUNT_SID`

## Run

```shell
./recreation-gob-notify
```
