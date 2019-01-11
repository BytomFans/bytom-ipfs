apackage providers

import (
	
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"
)

type HTTPProvider struct {
	address string
	timeout int32
	secure  bool
	client  *http.Client
}

var (
	httpClient *http.Client
)

func NewHTTPProvider(address string, timeout int32, secure bool) *HTTPProvider {
	return NewHTTPProviderWithClient(address, timeout, secure, httpClient)
}

func NewHTTPProviderWithClient(address string, timeout int32, secure bool, client *http.Client) *HTTPProvider {
	provider := new(HTTPProvider)
	provider.address = address
	provider.timeout = timeout
	provider.secure = secure
	provider.client = client

	return provider
}

func (provider HTTPProvider) SendRequest(v interface{}, method string, params interface{}) error {

	bodyString := util.JSONRPCObject{Params: params}

	prefix := "http://"
	if provider.secure {
		prefix = "https://"
	}

	body := strings.NewReader(bodyString.AsJsonString())
	req, err := http.NewRequest("POST", prefix+provider.address, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := provider.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, v)

}

func (provider HTTPProvider) Close() error { return nil }
