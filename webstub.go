package webstub

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Enable() {
	http.DefaultTransport = transport
}

func Disable() {
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}
}

type Request struct {
	method     string
	url        string
	statusCode int
	headers    map[string]string
	response   string
	file       string
}

func Register(rp Request) {
	key := rp.method + "-" + rp.url
	transport.stubs[key] = rp
}

type stubTransport struct {
	stubs map[string]Request
	http.Transport
}

func (m *stubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + "-" + req.URL.String()

	if stub_request, ok := m.stubs[key]; ok {
		if stub_request.file != "" {
			file, err := os.Open(stub_request.file)
			if err != nil {
				log.Fatal(err)
			}
			reader := bufio.NewReader(file)

			return http.ReadResponse(reader, req)
		}

		response := new(http.Response)

		if stub_request.statusCode > 0 {
			response.StatusCode = stub_request.statusCode
		} else {
			response.StatusCode = 200
		}

		response.Header = make(http.Header)
		for name, value := range stub_request.headers {
			response.Header.Add(name, value)
		}

		b := bytes.NewBufferString(stub_request.response)
		response.Body = ioutil.NopCloser(b)

		return response, nil
	}

	return m.Transport.RoundTrip(req)
}

func newStubTransport() *stubTransport {
	return &stubTransport{
		stubs: make(map[string]Request),
	}
}

var transport = newStubTransport()
