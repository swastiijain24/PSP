package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

type NpciClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewNpciClient(url string) *NpciClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, //must set to false in production
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &NpciClient{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: transport,
		},
	}
}


func (c *NpciClient) LinkAccount(ctx context.Context, vpaId string, accountId string, bankCode string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"vpa_id": vpaId,
		"account_id":accountId,
		"bank_code":bankCode,
	})
	
	url := "/vpa/register"
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json") 

	_, err := c.HTTPClient.Do(req)
	if err != nil {
		return err 
	}
	
	return nil 
}