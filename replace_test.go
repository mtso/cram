package main

import (
	"bytes"
	"reflect"
	"testing"
	"io/ioutil"
)

const (
	sample1 = "12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34\n12\n34"
	sample2 = "12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34\t12\t34"
)

func TestReplace(t *testing.T) {
	og := []byte(sample1)
	cram := Replace(og, newline, tab)
	uncram := Replace(cram, tab, newline)

	isReplaced := bytes.Index(cram, []byte{newline}) < 0
	if !isReplaced {
		t.Errorf("expected to not have newlines in: %s", cram)
	}

	if !reflect.DeepEqual(og, uncram) {
		t.Errorf("not equal\nOG: %s\nUncrammed: %s", og, uncram)
	}
}

func BenchmarkReplace(b *testing.B) {
	samp := []byte(sample1)
	for n := 0; n < b.N; n++ {
		Replace(fn, newline, tab)
	}
}

func BenchmarkStdReplace(b *testing.B) {
	samp := []byte(sample1)
	for n := 0; n < b.N; n++ {
		bytes.Replace(fn, []byte{newline}, []byte{tab}, -1)
	}
}
