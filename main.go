package crtsh

import (
	"crypto/x509"
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

const BASE_URL = "https://crt.sh/atom"

type CertificateEntry struct {
	ID          string
	Certificate *x509.Certificate
}

func Fetch(domain, exclude string) ([]CertificateEntry, error) {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, errors.Join(ErrorParseURL, err)
	}

	query := u.Query()
	query.Add("q", domain)
	query.Add("exclude", domain)
	u.RawQuery = query.Encode()

	feed, err := gofeed.NewParser().ParseURL(u.String())
	if err != nil {
		return nil, errors.Join(ErrorFetchRSS, err)
	}

	var entries []CertificateEntry
	for _, item := range feed.Items {
		desc := strings.NewReader(item.Description)
		node, err := html.Parse(desc)
		if err != nil {
			return nil, errors.Join(ErrorParseHTML, err)
		}

		first := node.LastChild.LastChild.LastChild.FirstChild

		s := parseHTMLElement(first)
		c, err := ParseCertificate(s)
		if err != nil {
			return nil, err
		}

		entries = append(entries, CertificateEntry{
			ID:          item.GUID,
			Certificate: c,
		})
	}

	return entries, nil
}
