package telex

import (
	"errors"
	"fmt"
)

type APMOptions struct {
	WebhookURL        string
	RequestTimeout    int
	Async             bool
	EnableTracing     bool
	TracingSampleRate float64
}

type Client struct {
	Options APMOptions
}

func Init(options APMOptions) (*Client, error) {
	if options.WebhookURL == "" {
		return nil, errors.New("APM initialization failed: WebhookURL is required")
	}

	client := &Client{Options: options}
	fmt.Printf("Telex APM initializd with webhook URL: %s\n", options.WebhookURL)
	return client, nil
}
