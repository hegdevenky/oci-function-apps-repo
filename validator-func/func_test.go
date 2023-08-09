package main

import (
	"bytes"
	"context"
	fdk "github.com/fnproject/fdk-go"
	"strings"
	"testing"
)

var (
	in  = new(bytes.Buffer)
	out = new(bytes.Buffer)
)

func TestMyHandler(t *testing.T) {
	handlerFunc := fdk.HandlerFunc(myHandler)

	// with input
	in.Reset()
	out.Reset()
	in.WriteString(`{"input": "hello"}`)
	handlerFunc.Serve(context.Background(), in, out)
	if strings.TrimSuffix(out.String(), "\n") != "true" {
		t.Errorf("assertion failure: expected %q but got %q", "true", out.String())
	}

	// with TimeoutInSeconds
	in.Reset()
	out.Reset()
	in.WriteString(`{"input": "hello", "executionTimeInSeconds": 10}`)
	handlerFunc.Serve(context.Background(), in, out)
	if strings.TrimSuffix(out.String(), "\n") != "true" {
		t.Errorf("assertion failure: expected %q but got %q", "true", out.String())
	}

	// without input
	in.Reset()
	out.Reset()
	handlerFunc.Serve(context.Background(), in, out)
	if strings.TrimSuffix(out.String(), "\n") != "false" {
		t.Errorf("assertion failure: expected %q but got %q", "false", out.String())
	}
}
