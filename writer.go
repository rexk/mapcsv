// Copyright 2014 Myung Sun Kim. All rights reserved.
// Use of this source code is governed by a MIT License
// that can be found in the LICENSE file
package mapcsv

import (
	"encoding/csv"
	"sort"
)

// A Writer writes records to a CSV encoded file.
//
// A Writer operates exactly same as encoding/csv/Writer except that it takes
// map[string]string instead of []string
//
// It also takes Fields name as []string if one wants to write fields name
// at the top of the stream
type Writer struct {
	*csv.Writer
	Fields []string
	Line   int
}

// NewWriter returns a new Writer that writes to csv.Writer
func NewWriter(w *csv.Writer) *Writer {
	return &Writer{w, []string{}, 0}
}

// Writer writes a single CSV record to csv.Writer by converting given map
// into []string
func (w *Writer) Write(record map[string]string) (err error) {
	numKeys := len(record)
	var keys []string
	for key, _ := range record {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	numFields := len(w.Fields)
	arrRecords := make([]string, numKeys)
	// Recording fields at the very first line
	if w.Line == 0 && numFields != 0 {
		w.Line++
		return w.Writer.Write(w.Fields)
	}

	// There is no Designated field, so convert key value as integer and fill out out array
	if numFields == 0 {
		for i, val := range keys {
			arrRecords[i] = record[val]
		}
	} else {
		for i, val := range w.Fields {
			arrRecords[i] = record[val]
		}
	}
	w.Line++
	return w.Writer.Write(arrRecords)
}

// Error reports any error that has occurred during a previous Write or Flush.
func (w *Writer) Error() error {
	return w.Writer.Write(nil)
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occured during the Flush, call Error.
func (w *Writer) Flush() {
	w.Writer.Flush()
}

// WriteAll writes multiple CSV reports to csv.Writer an then calls Flush.
func (w *Writer) WriteAll(records []map[string]string) (err error) {
	for _, record := range records {
		err = w.Write(record)
		if err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}
