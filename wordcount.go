package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
)

type Record struct {
	Word  string
	Count uint
}

func PrintRecords(w io.Writer, records []Record) {
	for _, record := range records {
		fmt.Fprintf(w, "%s: %d\n", record.Word, record.Count)
	}
}

type Config struct {
	IgnorePunct bool
}

type WordCounter struct {
	transform func(string) string
}

func NewWordCounter(config Config) WordCounter {
	var transform func(string) string

	if config.IgnorePunct {
		transform = removePunct
	} else {
		transform = identity
	}

	wc := WordCounter{
		transform: transform,
	}

	return wc
}

func identity(s string) string {
	return s
}

func removePunct(s string) string {
	return strings.TrimFunc(s, unicode.IsPunct)
}

func (wc WordCounter) Count(input io.Reader) []Record {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	countMap := make(map[string]uint)

	for scanner.Scan() {
		word := wc.transform(scanner.Text())
		if len(word) == 0 {
			continue
		}

		if _, ok := countMap[word]; !ok {
			countMap[word] = 0
		}
		countMap[word]++
	}

	records := recordsFromCountMap(countMap)
	sortByFrequencyThenName(records)

	return records
}

func recordsFromCountMap(countMap map[string]uint) []Record {
	records := []Record{}
	for word, count := range countMap {
		record := Record{
			Word:  word,
			Count: count,
		}
		records = append(records, record)
	}
	return records
}

func sortByFrequencyThenName(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].Count == records[j].Count {
			return records[i].Word < records[j].Word
		}
		return records[i].Count > records[j].Count
	})
}
