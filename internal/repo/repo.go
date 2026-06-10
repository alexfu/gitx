package repo

import (
	"io"
	"net/http"
	"net/url"
)

func DownloadExtention(name string) ([]byte, error) {
	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
		Path:   name,
	}
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
