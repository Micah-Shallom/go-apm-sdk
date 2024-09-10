package telex

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (c *Client)SendMetricsToWebHook(webhookURL string, metrics APMMetrics, timeout time.Duration) {
	payload, jsonErr := json.Marshal(metrics)
	if jsonErr != nil {
		log.Printf("Error marshalling metrics payload: %v", jsonErr)
		return
	}

	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	req, err := http.NewRequest(
		"POST",
		webhookURL,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending metrics to webhook: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Metrics sent to webhook: %s", webhookURL)
}


// func (c *Client)SendMetricsToKafka(metrics Metrics, timeout time.Duration) {}

// func (c *Client)SendMetricsToS3(metrics Metrics, timeout time.Duration) {}

// func (c *Client)SendMetricsToSNS(metrics Metrics, timeout time.Duration) {}

// func (c *Client)SendMetricsToSQS(metrics Metrics, timeout time.Duration) {}

