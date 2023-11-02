package wordcount_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pcjun97/wordcount"
)

const (
	input = "foo bar bar foo! bar.. ,"
)

func TestServerCount(t *testing.T) {
	config := wordcount.Config{
		IgnorePunct: false,
	}
	server := wordcount.NewServer(config)
	request, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(input))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := response.Body.String()
	want := `bar: 2
,: 1
bar..: 1
foo: 1
foo!: 1
`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestServerCountIgnorePunct(t *testing.T) {
	config := wordcount.Config{
		IgnorePunct: true,
	}
	server := wordcount.NewServer(config)
	request, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(input))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := response.Body.String()
	want := `bar: 3
foo: 2
`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
