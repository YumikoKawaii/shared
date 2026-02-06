package echo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ChargeEndpoint    = "api/v1/charge"
	DischargeEndpoint = "api/v1/discharge"
)

type httpClient struct {
	endpoint string
	client   *http.Client
}

func newHTTPClient(endpoint string) (Client, error) {
	return &httpClient{
		endpoint: endpoint,
		client:   &http.Client{},
	}, nil
}

func (c *httpClient) Charge(ctx context.Context, request ChargeRequest) (ChargeResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return ChargeResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.endpoint, ChargeEndpoint), bytes.NewReader(body))
	if err != nil {
		return ChargeResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return ChargeResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ChargeResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response ChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ChargeResponse{}, err
	}

	return response, nil
}

func (c *httpClient) Discharge(ctx context.Context, request DischargeRequest) (DischargeResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return DischargeResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.endpoint, DischargeEndpoint), bytes.NewReader(body))
	if err != nil {
		return DischargeResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return DischargeResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DischargeResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response DischargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return DischargeResponse{}, err
	}

	return response, nil
}
