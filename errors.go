package crtsh

import "github.com/pkg/errors"

var (
	ErrorParseBaseUrl     = errors.Errorf("failed to parse base URL")
	ErrorReadResponseBody = errors.Errorf("failed to read response body")
	ErrorParseJson        = errors.Errorf("failed to parse json")
	ErrorFetchCrtsh       = errors.Errorf("failed to fetch crt.sh")
)
