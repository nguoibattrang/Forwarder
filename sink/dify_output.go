package sink

import (
	"bytes"
	"context"
	"forwarder/config"
	"log"
	"net/http"
	"time"
)

type DifyProducer struct {
	url    string
	secret string
}

func NewDifyProducer(config *config.SinkConfig) *DifyProducer {
	return &DifyProducer{url: config.URL, secret: config.SecretKey}
}

// Produce sends data to the HTTP API.
func (inst *DifyProducer) Produce(ctx context.Context, data string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, inst.url, bytes.NewBufferString(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API response status: %d", resp.StatusCode)
		return err
	}

	return nil
}
