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

func (c *NpciClient) CreateAccount(ctx context.Context, name string, phone string, mpinHash string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"name":name, 
		"phone":phone,
		"mpin_hash":mpinHash,
	})

	url := "/accounts/"
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json") 

	_, err := c.HTTPClient.Do(req)
	if err != nil {
		return err 
	}
	return nil 
}

func (c *NpciClient) PaymentRequest(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64, mpin string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"transaction_id":transactionId,
		"payer_vpa":payerVpa,
		"payee_vpa":payeeVpa,
		"amount":amount,
		"mpin":mpin,
	})

	url := "/npci/payment"
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json") 

	_, err := c.HTTPClient.Do(req)
	if err != nil {
		return err 
	}
	return nil 

}