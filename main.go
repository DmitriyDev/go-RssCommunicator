package RssCommunicator

import (
	"io/ioutil"
	"net/http"
)

type Communicator struct {
}

type CommunicationHeader struct {
	Key   string
	Value string
}

type CommunicationRequest struct {
	Url     string
	Headers []CommunicationHeader
}

func DefaultHeaders() []CommunicationHeader {
	return []CommunicationHeader{
		{"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		{"User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Safari/605.1.15"},
		{"Host", "example.com"},
		{"Accept-Language", "en-us"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Connection", "keep-alive"},
	}
}

func (rc Communicator) build(rq CommunicationRequest) (*http.Request, error) {

	req, err := http.NewRequest("GET", rq.Url, nil)

	if err != nil {
		return nil, err
	}

	for _, header := range rq.headers {
		req.Header.Add(header.key, header.value)
	}

	return req, nil
}

func (rc Communicator) Execute(rq CommunicationRequest) (string, error) {

	client := &http.Client{}

	req, err := rc.build(rq)

	response, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}
