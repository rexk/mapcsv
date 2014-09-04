// Copyright 2014 Myung Sun Kim. All rights reserved.
// Use of this source code is governed by a MIT License
// that can be found in the LICENSE file
package mapcsv

import (
	"encoding/csv"
	"sort"
)

type Writer struct {
	*csv.Writer
	Fields []string
	Line   int
}

func NewWriter(w *csv.Writer) *Writer {
	return &Writer{w, []string{}, 0}
}

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

func (w *Writer) Error() error {
	return w.Writer.Write(nil)
}

func (w *Writer) Flush() {
	w.Writer.Flush()
}

func (w *Writer) WriteAll(records []map[string]string) (err error) {
	for _, record := range records {
		err = w.Write(record)
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}
