package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(stringUtilHandler))
}

const (
	UpperCase = "UPPER"
	LowerCase = "LOWER"
	TitleCase = "TITLE"
)

type StringUtilRequest struct {
	InputString string `json:"inputString"`
	Operation   string `json:"operation"`
	DryRun      bool   `json:"dryRun,omitempty"`
}

func stringUtilHandler(ctx context.Context, in io.Reader, out io.Writer) {
	request := new(StringUtilRequest)
	err := json.NewDecoder(in).Decode(request)
	if err != nil {
		json.NewEncoder(out).Encode(fmt.Sprintf("error: failed to parse request.\n "+
			"error message - %s", err.Error()))
		return
	}
	log.Printf("received input %+v", &request)
	result, err := stringUtil(request.InputString, request.Operation)

	if request.DryRun {
		log.Println("dryrun...")
		if err != nil {
			json.NewEncoder(out).Encode(false)
			return
		}
		json.NewEncoder(out).Encode(true)
		return
	}

	if err != nil {
		json.NewEncoder(out).Encode(fmt.Sprintf("error: failed to process request.\n "+
			"error message - %s", err.Error()))
		return
	}
	json.NewEncoder(out).Encode(result)
}

func stringUtil(input, op string) (string, error) {
	if input = strings.TrimSpace(input); input == "" {
		return "", fmt.Errorf("invalid input: input string is blank\n")
	}
	switch op = strings.TrimSpace(op); op != "" {
	case op == LowerCase:
		return strings.ToLower(input), nil
	case op == UpperCase:
		return strings.ToUpper(input), nil
	case op == TitleCase:
		return strings.Title(input), nil
	default:
		return "", fmt.Errorf("invalid input: invalid operation %q was supplied. "+
			"valid value are [%s,%s,%s]\n", op, UpperCase, LowerCase, TitleCase)
	}
}
