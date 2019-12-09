package dana

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Logger           *log.Logger
	SignatureEnabled bool
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	return Client{
		// LogLevel is the logging level used by the Dana library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		LogLevel: 2,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		SignatureEnabled: true,
	}
}

// ===================== HTTP CLIENT ================================================
var defHTTPTimeout = 15 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Request creation failed: ", err)
		}
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
	logLevel := c.LogLevel
	logger := c.Logger

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	start := time.Now()
	res, err := httpClient.Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("Dana response: ", string(resBody))
	}

	if v != nil && res.StatusCode == 200 {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}

		if c.SignatureEnabled {
			response := gjson.Get(string(resBody), "response")
			signature := gjson.Get(string(resBody), "signature")

			err := verifySignature(response.String(), signature.String(), c.PublicKey)
			if err != nil {
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

