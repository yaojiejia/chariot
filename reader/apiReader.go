package Reader

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yaojiejia/chariot/lib"
)

type APIReader struct {
	APIURL string
	APIKey string
	Cache  *lib.Cache
}

func NewAPIReader(apiURL, apiKey string) *APIReader {
	return &APIReader{
		APIURL: apiURL,
		APIKey: apiKey,
		Cache:  lib.NewCache(),
	}
}

func (a *APIReader) ReadAndCache() (lib.Cache, error) {
	// Read data from API
	data, err := a.Read()
	if err != nil {
		return lib.Cache{}, err
	}

	// Parse JSON data
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return lib.Cache{}, err
	}

	// Store each object in Redis cache
	for key, value := range jsonData {
		// Convert value back to JSON string
		valueJSON, err := json.Marshal(value)
		if err != nil {
			return lib.Cache{}, err
		}

		// Store in Redis using the object's key
		err = a.Cache.Set(key, string(valueJSON))
		if err != nil {
			return lib.Cache{}, err
		}
	}

	return *a.Cache, nil
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

func (a *APIReader) Get(key string) (string, error) {
	return a.Cache.Get(key)
}
