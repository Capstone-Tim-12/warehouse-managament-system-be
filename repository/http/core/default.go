package core

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/fadilahonespot/library/rest"
)


type wrapper struct {
	client rest.RestClient
}

var address string

func NewWrapper() CoreWrapper {
	restOptions := rest.Options{
		Address: os.Getenv("CORE_HOST"),
		Timeout: time.Duration(3 * time.Minute),
		SkipTLS: false,
	}
	client := rest.New(restOptions)
	address = restOptions.Address

	return &wrapper{client: client}
}

func getRequestHeaders(ctx context.Context) (headers http.Header) {
	headers = http.Header{
		"Content-Type": []string{"application/json"},
	}

	return
}

