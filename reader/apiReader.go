package Reader

import (
	"io"
	"net/http"
)

type APIReader struct {
	APIURL string
	APIKey string
	Cache  *Cache
}

func NewAPIReader(apiURL, apiKey string) *APIReader {
	return &APIReader{
		APIURL: apiURL,
		APIKey: apiKey,
	}
}

func (a *APIReader) Read() (string, error) {
	res, err := http.Get(a.APIURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resData), nil

}

func (a *APIReader) Connect() (error, string) {
	data, err := a.Read()
	if err != nil {
		return err, err.Error()
	}

	a.Cache = NewCache([][]string{{data}})
	return nil, "stored to the cache!"
}
