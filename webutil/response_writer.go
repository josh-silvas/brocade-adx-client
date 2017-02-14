package webutil

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

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

//returns the resp map as a json string with the error code provided
func ErrHandler(w http.ResponseWriter, r *http.Request, resp Response, code int) {
	http.Error(w, resp.String(), code)
	return
}

func JSONErrHandler(w http.ResponseWriter, r *http.Request, resp Response, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, resp)
}

//returns the resp map as a json string with a 200 OK
func ResHandler(w http.ResponseWriter, r *http.Request, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, resp.String())
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
