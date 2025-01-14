package sink

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nguoibattrang/forwarder/config"
)

type DifyProducer struct {
	hostname  string
	secret    string
	datasetId string
}

type ProcessRule struct {
	Mode string `json:"mode"`
}

type RequestPayload struct {
	Name              string      `json:"name"`
	Text              string      `json:"text"`
	IndexingTechnique string      `json:"indexing_technique"`
	ProcessRule       ProcessRule `json:"process_rule"`
}

func NewDifyProducer(config *config.SinkConfig) *DifyProducer {
	return &DifyProducer{hostname: config.Hostname, secret: config.SecretKey, datasetId: config.DatasetId}
}

// Produce sends data to the HTTP API.
func (inst *DifyProducer) Produce(name, data string) error {
	return createDifyKnowledgeDoc(inst.secret, inst.hostname, name, inst.datasetId, data)
}

func createDifyKnowledgeDoc(apiKey, hostname, name, datasetId, text string) error {
	url := fmt.Sprintf("https://%s/v1/datasets/%s/document/create_by_text", hostname, datasetId)

	payload := RequestPayload{
		Name:              name,
		Text:              text,
		IndexingTechnique: "high_quality",
		ProcessRule:       ProcessRule{Mode: "automatic"},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API %s responded with status %d: %s", url, resp.StatusCode, string(body))
	}

	return nil
}
