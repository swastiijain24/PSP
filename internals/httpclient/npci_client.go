package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client interface {
	DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error)
	LinkAccount(ctx context.Context, vpaId string, accountId string, bankCode string) error
	SetMpin(ctx context.Context, vpaId string, mpinEn string) error
	ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string, newMpinEn string) error
	GetBalance(ctx context.Context, vpaId string, mpinEn string) (int64, error)
	PaymentRequest(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64, mpin string) error
	GetStatus(ctx context.Context, transactionid string) (string, error)
}

type NpciClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewNpciClient(url string) Client {
	allowInsecureTLS := os.Getenv("ALLOW_INSECURE_TLS") == "true"
	tlsConfig := &tls.Config{
		InsecureSkipVerify: allowInsecureTLS,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &NpciClient{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
	}
}

func (c *NpciClient) DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"phone":     phone,
		"bank_code": bankCode,
	})

	url := fmt.Sprintf("%s/account/discover", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	var result []string
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (c *NpciClient) LinkAccount(ctx context.Context, vpaId string, accountId string, bankCode string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"vpa_id":     vpaId,
		"account_id": accountId,
		"bank_code":  bankCode,
	})

	url := fmt.Sprintf("%s/vpa/register", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	return nil
}

func (c *NpciClient) SetMpin(ctx context.Context, vpaId string, mpinEn string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"vpa_id":         vpaId,
		"mpin_encrypted": mpinEn,
	})

	url := fmt.Sprintf("%s/mpin", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}
	return nil
}

func (c *NpciClient) ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string, newMpinEn string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"vpa_id":             vpaId,
		"old_mpin_encrypted": oldMpinEn,
		"new_mpin_encrypted": newMpinEn,
	})

	url := fmt.Sprintf("%s/mpin", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	return nil
}

func (c *NpciClient) GetBalance(ctx context.Context, vpaId string, mpinEn string) (int64, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"vpa_id":         vpaId,
		"mpin_encrypted": mpinEn,
	})

	url := fmt.Sprintf("%s/mpin", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	var result int64
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (c *NpciClient) PaymentRequest(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64, mpin string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"transaction_id": transactionId,
		"payer_vpa":      payerVpa,
		"payee_vpa":      payeeVpa,
		"amount":         amount,
		"mpin":           mpin,
	})

	url := fmt.Sprintf("%s/npci/payment", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}
	return nil

}

func (c *NpciClient) GetStatus(ctx context.Context, transactionid string) (string, error) {
	url := fmt.Sprintf("%s/npci/status/%s", c.BaseURL, transactionid)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("X-API-Key", os.Getenv("PSP_API_KEY"))
	req.Header.Set("X-PSP-ID", os.Getenv("PSP_ID"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "ERROR", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "ERROR", err
	}

	var result string
	json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
