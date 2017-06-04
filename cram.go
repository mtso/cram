package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// - read file
// --- replace newlines with tabs?
// - write file and append .cram
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: needs a filepath")
		os.Exit(0)
	}
	name := os.Args[1]

	og, err := ioutil.ReadFile(name)
	check(err)

	err = os.Remove(name)
	check(err)

	if bytes.Count(og, []byte{'\n'}) < 1 {
		d := uncram(og)
		err = ioutil.WriteFile(strings.Replace(name, ".cram", "", -1), d, os.ModePerm)
	} else {
		d := cram(og)
		err = ioutil.WriteFile(name+".cram", d, os.ModePerm)
	}
	check(err)
}

func cram(og []byte) []byte {
	return bytes.Replace(og, []byte{'\n'}, []byte{'\t'}, -1)
}

func uncram(og []byte) []byte {
	return bytes.Replace(og, []byte{'\t'}, []byte{'\n'}, -1)
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
