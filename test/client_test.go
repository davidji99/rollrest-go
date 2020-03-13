package test

import (
	"fmt"
	"github.com/davidji99/rollbar-go/rollbar"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux       *http.ServeMux
	server    *httptest.Server
	client    *rollbar.Client
	clientErr error
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	fmt.Println(server.URL)
	client, clientErr = rollbar.New(rollbar.BaseURL(server.URL), rollbar.AuthAAT("account_access_token"),
		rollbar.AuthPAT("project_access_token"))
	if clientErr != nil {
		panic(clientErr)
	}

	return func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("../testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
