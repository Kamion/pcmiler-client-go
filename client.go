package pcmiler

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	APIKey string
}

// New instantiates new PCMiler client with your API key and optional dataset, defaults to Current
func New(key string) *Client {
	return &Client{
		APIKey: key,
	}
}

func (c *Client) request(method, url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, host+url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// todo: hide debug information once I'm done
	// dump, _ := httputil.DumpRequest(req, true)
	// log.Println(string(dump))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// dump, _ := httputil.DumpResponse(res, true)
	// log.Println(string(dump))

	return ioutil.ReadAll(res.Body)
}
