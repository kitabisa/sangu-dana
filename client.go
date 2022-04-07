package dana

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"moul.io/http2curl"
	"net/http"
	"strings"
	"time"
)

// Client struct
type Client struct {
	BaseUrl          string
	Version          string
	ClientId         string
	ClientSecret     string
	PrivateKey       []byte
	PublicKey        []byte
	LogLevel         int
	Logger           Logger
	SignatureEnabled bool
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
	}

	logger := *NewLogger(logOption)

	return Client{
		// LogLevel is the logging level used by the Dana library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		LogLevel:         2,
		Logger:           logger,
		SignatureEnabled: true,
	}
}

// ===================== HTTP CLIENT ================================================
var defHTTPTimeout = 15 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		c.Logger.Info("Request creation failed: %v ", err)
		return nil, err
	}

	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}

	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request, v interface{}) error {
	c.Logger.Info("Start requesting: %v ", req.URL)

	command, _ := http2curl.GetCurlCommand(req)
	start := time.Now()
	res, err := httpClient.Do(req)
	if err != nil {
		c.Logger.Error("Request failed. Error : %v , Curl Request : %v", err, command)
		return err
	}
	defer res.Body.Close()

	c.Logger.Info("Completed in %v", time.Since(start))
	c.Logger.Info("Curl Request: %v ", command)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error("Cannot read response body: %v ", err)
		return err
	}

	c.Logger.Info("DANA HTTP status response : %d", res.StatusCode)
	c.Logger.Info("DANA response body : %s", string(resBody))

	if v != nil && res.StatusCode == 200 {
		if err = json.Unmarshal(resBody, v); err != nil {
			c.Logger.Error("Failed unmarshal body: %v ", err)
			return err
		}

		// Dana endpoint V1 doesn't return signature in response, so we don't need to verify signature again
		if strings.Contains(req.URL.String(), "/v1/") {
			c.Logger.Info("Req URL Contains Dana endpoint V1")
			return nil
		}

		if c.SignatureEnabled {
			response := gjson.Get(string(resBody), "response")
			signature := gjson.Get(string(resBody), "signature")

			err = verifySignature(response.String(), signature.String(), c.PublicKey)
			if err != nil {
				c.Logger.Error("verifySignature failed: %v ", err)
				return err
			}
		}
	}

	return nil
}

// Call the Dana API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, header, body)
	if err != nil {
		return err
	}

	return c.ExecuteRequest(req, v)
}

// ===================== END HTTP CLIENT ================================================
