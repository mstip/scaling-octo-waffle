package pushover

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Pushover struct {
	Token string
	User  string
}

func NewPushover(token string, user string) *Pushover {
	return &Pushover{
		Token: token,
		User:  user,
	}
}

func (p *Pushover) SendMessage(text string) error {
	resp, err := http.PostForm("https://api.pushover.net/1/messages.json", url.Values{
		"token":   {p.Token},
		"user":    {p.User},
		"message": {text}})
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code was: %d", resp.StatusCode))
	}

	return nil
}
