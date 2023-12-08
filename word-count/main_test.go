package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n\n\n\n")
	expected := 4
	res := count(b, false, false)
	if res != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, res)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n\n\n\n")
	expected := 4
	res := count(b, true, false)
	if res != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, res)
	}
}

func TestByteCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")
	expected := 23
	res := count(b, false, true)
	if res != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, res)
	}
}
