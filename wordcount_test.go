package wordcount_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/pcjun97/wordcount"
)

var (
	input1 = "foo bar bar"
	input2 = strings.ReplaceAll(input1, " ", "\t")
	input3 = "foo. _bar!! foo, _ bar bar"
	input4 = `foo bar
foo. _bar!! @foo foo,
foo bar bar`
)

func TestCount(t *testing.T) {
	cases := []struct {
		Name  string
		Input string
		Want  []wordcount.Record
	}{
		{"input1", input1, []wordcount.Record{{"bar", 2}, {"foo", 1}}},
		{"input2", input2, []wordcount.Record{{"bar", 2}, {"foo", 1}}},
		{"input3", input3, []wordcount.Record{{"bar", 2}, {"_", 1}, {"_bar!!", 1}, {"foo,", 1}, {"foo.", 1}}},
		{"input4", input4, []wordcount.Record{{"bar", 3}, {"foo", 2}, {"@foo", 1}, {"_bar!!", 1}, {"foo,", 1}, {"foo.", 1}}},
	}

	config := wordcount.Config{
		IgnorePunct: false,
	}
	wc := wordcount.NewWordCounter(config)

	for _, c := range cases {
		t.Run("count "+c.Name, func(t *testing.T) {
			got := wc.Count(strings.NewReader(c.Input))

			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("got %v, want %v", got, c.Want)
			}
		})
	}
}

func TestCountIgnorePunct(t *testing.T) {
	cases := []struct {
		Name  string
		Input string
		Want  []wordcount.Record
	}{
		{"input1", input1, []wordcount.Record{{"bar", 2}, {"foo", 1}}},
		{"input2", input2, []wordcount.Record{{"bar", 2}, {"foo", 1}}},
		{"input3", input3, []wordcount.Record{{"bar", 3}, {"foo", 2}}},
		{"input4", input4, []wordcount.Record{{"foo", 5}, {"bar", 4}}},
	}

	config := wordcount.Config{
		IgnorePunct: true,
	}
	wc := wordcount.NewWordCounter(config)

	for _, c := range cases {
		t.Run("count ignore punctuation "+c.Name, func(t *testing.T) {
			got := wc.Count(strings.NewReader(c.Input))

			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("got %v, want %v", got, c.Want)
			}
		})
	}
}

func TestPrintRecords(t *testing.T) {
	records := []wordcount.Record{
		{"foo", 24},
		{"bar", 17},
		{"baz", 3},
	}

	buffer := bytes.Buffer{}
	wordcount.PrintRecords(&buffer, records)

	got := buffer.String()
	want := `foo: 24
bar: 17
baz: 3
`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
