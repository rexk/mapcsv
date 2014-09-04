// Copyright 2014 Myung Sun Kim. All rights reserved.
// Use of this source code is governed by a MIT License
// that can be found in the LICENSE file

package mapcsv

import (
	"encoding/csv"
	"io"
	"strconv"
)

// A Reader read record form a CSV-encoded file.
//
// This reader returns map[string]string rather than []string as a result
//
// If FieldsAtFirstLine is true, csv record at first line are used as keys for map[string]string.
// Otherwise, key for map will be a string representation of column number.
//
//       { "0": "1", "1": "2" }
//
type Reader struct {
	*csv.Reader
	FieldsAtFirstLine bool
	Line              int
	fields            []string
}

// NewReader returns a new Reader that wrapps original encoding/csv reader
func NewReader(reader *csv.Reader) *Reader {
	return &Reader{reader, false, 0, make([]string, 1)}
}

// Read reads and decodes one record from csv.Reader. The record is a map[string]string
// with key as a number or name field depending on FieldsAtFirstLine value
func (r *Reader) Read() (map[string]string, error) {
	record, err := r.Reader.Read()
	if err != nil {
		return nil, err
	}
	if r.Line == 0 && r.FieldsAtFirstLine {
		r.fields = record
		r.Line++
		return r.Read()
	}
	if r.Line == 0 {
		fieldLen := len(record)
		var fields []string
		for i := 0; i < fieldLen; i++ {
			fields = append(fields, strconv.Itoa(i))
		}
		r.fields = fields
		r.Line++
	}
	r.Line++
	recordMap := make(map[string]string)
	for index, val := range record {
		key := r.fields[index]
		recordMap[key] = val
	}

	return recordMap, nil
}

// ReadAll reads all the remaning records from
// Each record is a slice of map[string]string
// A sucessful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reportedg
func (r *Reader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}
