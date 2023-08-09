package main

import (
	"context"
	"encoding/json"
	fdk "github.com/fnproject/fdk-go"
	"io"
	"log"
	"time"
)

const defaultTimeoutInSeconds = 65

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type FnIO struct {
	Input            string `json:"input,omitempty"`
	TimeoutInSeconds int    `json:"timeoutInSeconds,omitempty"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	log.Print("INFO: inside validator function")
	ip := new(FnIO)
	err := json.NewDecoder(in).Decode(ip)
	if err != nil {
		log.Printf("ERROR: failed to parse the request %+v. "+
			"proceeding with default values\n", *ip)
	}
	log.Printf("INFO: received input %+v\n", *ip)
	timeout := defaultTimeoutInSeconds
	if ip.TimeoutInSeconds != 0 {
		timeout = ip.TimeoutInSeconds
	}
	log.Printf("INFO: executing business logic...time remaining %ds\n", timeout)
	time.Sleep(time.Duration(timeout) * time.Second)
	if ip.Input != "" || ip.TimeoutInSeconds != 0 {
		json.NewEncoder(out).Encode(true)
		return
	}
	log.Print("INFO: didn't receive an input")
	json.NewEncoder(out).Encode(false)
}
