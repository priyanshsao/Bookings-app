package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"ab", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"search-a", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-r", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-a", "/search-availability", "POST", []postData{
		{key: "start", value: "2024-03-23"},
		{key: "end", value: "2024-03-29"},
	}, http.StatusOK},
	{"post-search-a-j", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2024-03-23"},
		{key: "end", value: "2024-03-29"},
	}, http.StatusOK},
	{"post-make-r", "/make-reservation", "POST", []postData{
		{key: "first-name", value: "John"},
		{key: "last-name", value: "Doe"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expeted %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {

			values := url.Values{}

			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expeted %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
