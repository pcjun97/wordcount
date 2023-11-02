package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pcjun97/wordcount"
	flag "github.com/spf13/pflag"
)

var ignorePunct bool
var server bool
var port uint

func init() {
	flag.BoolVarP(&ignorePunct, "ignore-punct", "i", false, "ignore punctuation when count words")
	flag.BoolVarP(&server, "server", "s", false, "run as a server")
	flag.UintVarP(&port, "port", "p", 3000, "port to listen")
	flag.Parse()

	argsLen := len(flag.Args())
	if argsLen > 1 || (!server && argsLen != 1) {
		fmt.Fprintln(os.Stderr, "usage: wordcount [-i] [-s] [-p port] [file]")
		os.Exit(1)
	}

	if server && argsLen > 0 {
		fmt.Fprintln(os.Stderr, "usage: wordcount -s [-i] [-p port]")
		os.Exit(1)
	}
}

func main() {
	if server {
		runServer()
	} else {
		runOnce()
	}
}

func runOnce() {
	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v", err)
	}
	defer f.Close()

	config := wordcount.Config{
		IgnorePunct: ignorePunct,
	}
	wc := wordcount.NewWordCounter(config)

	var records []wordcount.Record
	records = wc.Count(f)
	wordcount.PrintRecords(os.Stdout, records)
}

func runServer() {
	config := wordcount.Config{
		IgnorePunct: ignorePunct,
	}
	server := wordcount.NewServer(config)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
