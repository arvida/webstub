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
	Method     string
	Url        string
	StatusCode int
	Headers    map[string]string
	Response   string
	File       string
}

func Register(rp Request) {
	key := rp.Method + "-" + rp.Url
	transport.stubs[key] = rp
}

type stubTransport struct {
	stubs map[string]Request
	http.Transport
}

func (m *stubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + "-" + req.URL.String()

	if stub_request, ok := m.stubs[key]; ok {
		if stub_request.File != "" {
			file, err := os.Open(stub_request.File)
			if err != nil {
				log.Fatal(err)
			}
			reader := bufio.NewReader(file)

			return http.ReadResponse(reader, req)
		}

		response := new(http.Response)

		if stub_request.StatusCode > 0 {
			response.StatusCode = stub_request.StatusCode
		} else {
			response.StatusCode = 200
		}

		response.Header = make(http.Header)
		for name, value := range stub_request.Headers {
			response.Header.Add(name, value)
		}

		b := bytes.NewBufferString(stub_request.Response)
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
