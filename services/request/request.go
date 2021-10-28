package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func New(method string, url string, bodyMap map[string]string) (*http.Request, error) {

	var body io.Reader

	if len(bodyMap) > 0 {
		jsonBody, _ := json.Marshal(bodyMap)
		body = bytes.NewBuffer(jsonBody)
	} else {
		body = nil
	}

	req, err := http.NewRequest(method, url, body)

	return req, err

}

func AddHeaders(req *http.Request, k, v string) {

	req.Header.Add(k, v)
}

func DisableSSLTLSValidation() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func Do(req *http.Request) string {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	//log.Printf(sb)
	return sb
}
