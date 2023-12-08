package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T){
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	res := count(b)
	if res != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, res)
	}
}