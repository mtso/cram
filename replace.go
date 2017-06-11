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
	og []byte
	new []byte
	from byte
	to byte
}

func ReplaceAsync(done chan bool, buf Buf, start int, end int) {
	if end - start < 1 {
		buf.Lock()
		defer buf.Unlock()
		
		if buf.og[start] == buf.from {
			buf.new[start] = buf.to
		} else {
			buf.new[start] = buf.og[start]
		}
		done <- true
		return
	}

	// 0 1 2 3 4 5
	mid := start + (end - start) / 2
	ReplaceAsync(done, buf, start, mid)
	ReplaceAsync(done, buf, mid+1, end)
}

func Replace(raw []byte, from byte, to byte) []byte {
	new := make([]byte, len(raw))
	buf := Buf{
		&sync.Mutex{},
		raw,
		new,
		from,
		to,
	}
	done := make(chan bool)
	go ReplaceAsync(done, buf, 0, len(raw)-1)

	for i := 0; i < len(raw); i++ {
		<-done
	}
	return buf.new
}

func main() {
	f, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	Replace(f, newline, tab)

	fmt.Printf("%s\n", f)
}
