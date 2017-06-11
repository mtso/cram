package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str := []byte("")
	for i := 0; i < 1000000; i++ {
		if i % 2 == 0 {
			str = append(str, []byte("ba\n")...)
			// str += "ba\n"
			continue
		}
		str = append(str, []byte("foo\n")...)
		// str += "foo\n"
	}
	err := ioutil.WriteFile("long.txt", str, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
