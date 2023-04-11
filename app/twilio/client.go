package twilio

import "github.com/sfreiberg/gotwilio"

type Twilio struct {
	AccountSid string
	AuthToken  string
	From       string
	To         string
	client     gotwilio.Twilio
}

func NewTwilio(accountSid, authToken, from, to string) *Twilio {
	return &Twilio{
		AccountSid: accountSid,
		AuthToken:  authToken,
		From:       from,
		To:         to,
		client:     *gotwilio.NewTwilioClient(accountSid, authToken),
	}
}

func (t *Twilio) SendSMS(text string) (*gotwilio.SmsResponse, error, error) {
	return t.client.SendSMS(t.From, t.To, text, "", "")
}
