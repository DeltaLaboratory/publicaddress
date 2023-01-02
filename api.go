package publicaddress

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Information struct {
	// IP represents IP address of machine, can be IPv4 or IPv6
	IP string `json:"ip"`
	// Country is country based on machine ip, two character uppercase country-code
	Country string `json:"country"`
}

func GetInformation() (*Information, error) {
	response, err := http.Get("https://public-address.api.deltalab.dev")
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to perform request: bad status(%d)", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var result Information
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}
