package soap

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"strings"
	"encoding/json"
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

//XMLPost will porform a post request with the data provided, it will add content-type header xml to the header map.
func XMLPost(payload []byte, path string, headers map[string]string, insecure bool) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/xml"
	return Post(payload, path, headers, insecure)
}

// generic response writter for APIs
type Response map[string]interface{}

//This method returns a xml marshaled response
func (r Response) XML() string {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return strings.Replace(string(b), "%", "%%", -1)
}

// XMLWithHeader will take a structure and xml marshal with
// the <?xml version="1.0" encoding="UTF-8"?> header prepended
// to the XML request.
func XMLMarshalHead(r interface{}) string {
	rv := []byte(xml.Header)
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	rv = append(rv, b...)
	return strings.Replace(string(rv), "\\\"", "\"", -1)
}

//This method returns a json marshaled response
func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return strings.Replace(string(b), "%", "%%", -1)
}

//returns the resp map as a xml string with the error code provided
func XMLErrHandler(w http.ResponseWriter, r *http.Request, resp Response, code int) {
	w.Header().Set("Content-Type", "application/xml")
	http.Error(w, resp.XML(), code)
	return
}

//returns the resp map as a xml string with a 200 OK
func XMLResHandler(w http.ResponseWriter, r *http.Request, resp Response) {
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, resp.XML())
	return
}