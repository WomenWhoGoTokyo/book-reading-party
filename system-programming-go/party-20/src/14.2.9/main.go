package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	emitter := make(chan interface{})
	source := observable.Observable(emitter)
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func ") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}
	
	sub := source.Subscribe(watcher)
	
	go func() {
		content, err := os.ReadFile("reactive.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()
	
	<-sub
}
