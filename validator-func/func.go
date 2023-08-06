package main

import (
	"context"
	"encoding/json"
	fdk "github.com/fnproject/fdk-go"
	"io"
	"log"
	"time"
)

const waitTimeInSeconds = 65

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type FnIO struct {
	Input string `json:"input,omitempty"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	log.Print("Inside validator function")
	ip := &FnIO{Input: ""}
	json.NewDecoder(in).Decode(ip)
	log.Printf("executing business logic...time remaining %ds\n", waitTimeInSeconds)
	time.Sleep(waitTimeInSeconds * time.Second)
	if ip.Input != "" {
		log.Printf("Received the input - %s\n", ip.Input)
		json.NewEncoder(out).Encode(true)
		return
	}
	log.Print("Didn't receive an input")
	json.NewEncoder(out).Encode(false)
}
