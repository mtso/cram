package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var files = []string{
	"./test_data/yarn.lock",
	"./test_data/package.json",
}

func TestIdentity(t *testing.T) {
	for i := 0; i < len(files); i++ {
		og, err := ioutil.ReadFile(files[i])
		if err != nil {
			t.Error(err)
			return
		}

		crammed := cram(og)
		uncrammed := uncram(crammed)

		if !reflect.DeepEqual(uncrammed, og) {
			t.Errorf("og != uncrammed result: %s", files[i])
		}
	}
}
