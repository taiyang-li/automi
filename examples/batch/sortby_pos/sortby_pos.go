package main

import (
	"fmt"

	"github.com/taiyang-li/automi/collectors"
	"github.com/taiyang-li/automi/emitters"
	"github.com/taiyang-li/automi/stream"
)

func main() {
	data := emitters.Slice([][]string{
		{"request", "/i/a", "00:11:51:AA", "accepted"},
		{"response", "/i/a/", "00:11:51:AA", "failed"},
		{"request", "/i/b", "00:11:22:33", "accepted"},
		{"response", "/i/b", "00:11:22:33", "served"},
		{"request", "/i/c", "00:11:51:AA", "accepted"},
		{"response", "/i/c", "00:11:51:AA", "served"},
		{"request", "/i/d", "00:BB:22:DD", "accepted"},
		{"response", "/i/d", "00:BB:22:DD", "failed"},
	})

	stream := stream.New(data)

	stream.Filter(func(e []string) bool {
		return (e[0] == "response")
	})

	// sort returns [][]string
	stream.Batch().SortByPos(3)

	stream.Into(collectors.Func(func(data interface{}) error {
		items := data.([][]string)
		for _, item := range items {
			fmt.Printf("%v\n", item)
		}
		return nil
	}))

	// open the stream
	if err := <-stream.Open(); err != nil {
		fmt.Println(err)
		return
	}
}
