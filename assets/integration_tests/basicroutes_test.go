// +build integration

package integration_tests

import (
	"net/http"
	"testing"
)


var mainRoute = "http://localhost:8080/"
var routes200 = []string{
	"",
	"health-check/",
	"static/",
	"static/css/main.css",
	"static/img/favicon.ico",
}

func TestRoute200(t *testing.T) {
	for _, route := range routes200 {
		resp, _ := http.Get(mainRoute + route)

		if resp.StatusCode != 200 {
			t.Error("Expected 200 response for", route)
		}
	}
}

func TestRoute404(t *testing.T) {
	resp, _ := http.Get(mainRoute + "something/that/should/not/exist/")

	if resp.StatusCode != 404 {
		t.Error("Expected 404 response")
	}
}