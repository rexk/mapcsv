// Copyright 2014 Myung Sun Kim. All rights reserved.
// Use of this source code is governed by a MIT License
// that can be found in the LICENSE file

// This test code is slight modification from encoding/csv/writer_test.go
package mapcsv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"testing"
)

var writeTests = []struct {
	Input   []map[string]string
	Output  string
	UseCRLF bool
}{
	{Input: []map[string]string{{"0": "abc"}}, Output: "abc\n"},
	{Input: []map[string]string{{"0": "abc"}}, Output: "abc\r\n", UseCRLF: true},
	{Input: []map[string]string{{"0": `"abc"`}}, Output: `"""abc"""` + "\n"},
	{Input: []map[string]string{{"0": `a"b`}}, Output: `"a""b"` + "\n"},
	{Input: []map[string]string{{"0": `"a"b"`}}, Output: `"""a""b"""` + "\n"},
	{Input: []map[string]string{{"0": " abc"}}, Output: `" abc"` + "\n"},
	{Input: []map[string]string{{"0": "abc,def"}}, Output: `"abc,def"` + "\n"},
	{Input: []map[string]string{{"0": "abc", "1": "def"}}, Output: "abc,def\n"},
	{Input: []map[string]string{{"0": "abc"}, {"0": "def"}}, Output: "abc\ndef\n"},
	{Input: []map[string]string{{"0": "abc\ndef"}}, Output: "\"abc\ndef\"\n"},
	{Input: []map[string]string{{"0": "abc\ndef"}}, Output: "\"abc\r\ndef\"\r\n", UseCRLF: true},
	{Input: []map[string]string{{"0": "abc\rdef"}}, Output: "\"abcdef\"\r\n", UseCRLF: true},
	{Input: []map[string]string{{"0": "abc\rdef"}}, Output: "\"abc\rdef\"\n", UseCRLF: false},
}

func TestWrite(t *testing.T) {
	for n, tt := range writeTests {
		b := &bytes.Buffer{}
		g := csv.NewWriter(b)
		g.UseCRLF = tt.UseCRLF
		f := NewWriter(g)
		err := f.WriteAll(tt.Input)
		if err != nil {
			t.Errorf("Unexpected error: %s\n", err)
		}
		out := b.String()
		if out != tt.Output {
			t.Errorf("#%d: out=%q want %q", n, out, tt.Output)
		}
	}
}

type errorWriter struct{}

func (e errorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Test")
}

func TestError(t *testing.T) {
	b := &bytes.Buffer{}
	g := csv.NewWriter(b)
	// Proxying CSV Writer with mapcsv Writer
	f := NewWriter(g)
	f.Write(map[string]string{"0": "abc"})
	f.Flush()
	err := f.Error()

	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}

	f = NewWriter(csv.NewWriter(errorWriter{}))
	f.Write(map[string]string{"0": "abc"})
	f.Flush()
	err = f.Error()

	if err == nil {
		t.Error("Error should not be nil")
	}
}
