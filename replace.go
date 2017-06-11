package main

import (
	"io/ioutil"
	"sync"
	"log"
	"fmt"
)

const (
	newline = '\n'
	tab = '\t'
)

type Buf struct {
	*sync.Mutex
	B []byte
}

func Replace(buf []byte, start int, end int) {
	if end - start < 1 {
		buf.Lock()
		defer buf.Unlock()
		
		if buf[start] == newline {
			buf[start] = tab
		}
		return
	}

	// 0 1 2 3 4 5
	mid := start + (end - start) / 2
	Replace(buf, start, mid)
	Replace(buf, mid+1, end)
}

func RecReplace(buf []byte, start int, end int) {

}

func ReplaceAll(buf []byte) {
	Replace(buf, 0, len(buf) - 1)
}

func main() {
	f, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// buf := Buf{
	// 	&sync.Mutex{},
	// 	f,
	// }
	// Replace(f, 0, len(f)-1)
	ReplaceAll(f)

	fmt.Printf("%s\n", f)
}
