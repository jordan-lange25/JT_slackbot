package chunky

import (
	"fmt"

	"github.com/slack-go/slack"
)

type Chunky struct {
	token         string
	signingSecret string
	client        *slack.Client
	opts          ChunkyOpts
}

type ChunkyOpts struct {
	debug bool
}

func (c *Chunky) InitalizeClient(token string) {

	c.token = token
	c.client = slack.New(token)
}

func (c *Chunky) SetSigningToken(secret string) {
	c.signingSecret = secret
}

func (c *Chunky) PostMessage(channel, message string) error {
	repsChannel, respTimestamp, err := c.client.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		if c.opts.debug {
			fmt.Println(err)
		}
		return err
	}

	if c.opts.debug {
		fmt.Printf("Channel Response: %s\n", repsChannel)
		fmt.Printf("Timestamp: %s\n", respTimestamp)
	}

	return nil
}
