package crtsh

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/cockroachdb/errors"
)

const BASE_URL = "https://crt.sh/json"

func Fetch(domain string) ([]Certificate, error) {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, errors.Join(ErrorParseBaseUrl, err)
	}

	query := u.Query()
	query.Add("q", domain)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())

	if err != nil {
		return nil, errors.Join(ErrorFetchCrtsh, err)
	}
	defer resp.Body.Close()

	var certs []Certificate
	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(ErrorReadResponseBody, err)
	}

	err = json.Unmarshal(byteArray, &certs)
	if err != nil {
		return nil, errors.Join(ErrorParseJson, err)
	}

	return certs, nil
}
