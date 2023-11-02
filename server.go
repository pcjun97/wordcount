package wordcount

import "net/http"

type Server struct {
	wc WordCounter
}

func NewServer(config Config) http.HandlerFunc {
	wc := NewWordCounter(config)
	return func(w http.ResponseWriter, r *http.Request) {
		records := wc.Count(r.Body)
		PrintRecords(w, records)
	}
}
