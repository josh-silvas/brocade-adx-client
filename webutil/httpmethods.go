package webutil

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

//BasicAuthGet adds basic auth and perform a request to path
func BasicAuthGet(path, username, password string) ([]byte, int, error) {
	request, err := BasicAuthRequest(path, username, password)
	if err != nil {
		return nil, 0, err
	}
	return GetRequest(request, true)
}

//BasicAuthRequest creates a "GET" request with basic auth and return it
func BasicAuthRequest(path, username, password string) (*http.Request, error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(username, password)
	if err != nil {
		return nil, err
	}
	return request, nil
}

//Get makes a GET request to path
func Get(path string) ([]byte, int, error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, 0, err
	}
	return GetRequest(request, true)
}

//GetRequest executes a request and returns an []byte response, response code, and error
func GetRequest(request *http.Request, insecure bool) ([]byte, int, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	return ResponseCheck(response)
}

//XMLBasicAuthGet adds basic auth and perform a request to path
func XMLBasicAuthGet(path, username, password string) ([]byte, int, error) {
	request, err := BasicAuthRequest(path, username, password)
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("accept", "application/xml")
	return GetRequest(request, true)
}

//XMLBasicAuthGet adds basic auth and perform a request to path
func XMLADXBasicAuthPost(path, method, payload, username, password string) ([]byte, int, error) {
	req, err := http.NewRequest("POST", path, bytes.NewBufferString(payload))
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("SOAPAction", fmt.Sprintf("\"urn:webservicesapi#%v\"", method))
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	return ResponseCheck(resp)
}

//XMLGet performs a "GET" request to the path and appends the application/xml
func XMLGet(path string) ([]byte, int, error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("accept", "application/xml")
	return Get(path)
}

//XMLGetRequest performs a get request and sets the application/xml header to accept
func XMLGetRequest(request *http.Request) ([]byte, int, error) {
	request.Header.Set("accept", "application/xml")
	return GetRequest(request, true)
}

//JSONAuthGet performs a request with basic auth and application/json
func JSONAuthGet(path, username, password string) ([]byte, int, error) {
	request, err := BasicAuthRequest(path, username, password)
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("accept", "application/json")
	return GetRequest(request, true)
}

//JSONGet performs a request with the application/json header
func JSONGet(path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.Header.Set("accept", "application/json")
	return GetRequest(request, insecure)
}

//JSONGetClose performs a request with application/json header and the connection close header
func JSONGetClose(path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.Header.Set("accept", "application/json")
	request.Header.Set("Connection", "close")
	return GetRequest(request, insecure)
}

//JSONGetRequest performs a get request and sets the application/json header to accept
func JSONGetRequest(request *http.Request) ([]byte, int, error) {
	request.Header.Set("accept", "application/json")
	return GetRequest(request, true)
}

//ResponseCheck reads the response and return data, status code or error it encountered.
func ResponseCheck(response *http.Response) ([]byte, int, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return body, response.StatusCode, nil
}

//Post will porform a post request with the data provided.
func Post(payload []byte, path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(payload))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	return ResponseCheck(resp)
}

//JSONPost will porform a post request with the data provided, it will add content-type header json to the header map.
func JSONPost(payload []byte, path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"
	return Post(payload, path, headers, insecure)
}

//XMLPost will porform a post request with the data provided, it will add content-type header xml to the header map.
func XMLPost(payload []byte, path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/xml"
	return Post(payload, path, headers, insecure)
}
