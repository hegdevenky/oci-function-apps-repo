package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type FnIO struct {
	Input  string `json:"input,omitempty"`
	Result string `json:"result,omitempty"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	log.Print("Inside validator function")
	ip := &FnIO{Input: ""}
	json.NewDecoder(in).Decode(ip)
	log.Println("executing business logic...time remaining 60s")
	time.Sleep(60 * time.Second)
	if ip.Input != "" {
		log.Printf("Received the input %s\n", ip.Input)
	}
	log.Print("Didn't receive an input")
	json.NewEncoder(out).Encode(true)
}
