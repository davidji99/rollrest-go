package rollbar

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	baseURLPath = "/api-v3"
)

func setup(client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	//apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	client = New()
}
