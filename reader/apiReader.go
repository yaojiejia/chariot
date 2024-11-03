package Reader

import (
	"encoding/json"
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
		Cache:  NewCache(),
	}
}

func (a *APIReader) ReadAndCache() error {
	// Read data from API
	data, err := a.Read()
	if err != nil {
		return err
	}

	// Parse JSON data
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return err
	}

	// Store each object in Redis cache
	for key, value := range jsonData {
		// Convert value back to JSON string
		valueJSON, err := json.Marshal(value)
		if err != nil {
			return err
		}

		// Store in Redis using the object's key
		err = a.Cache.Set(key, string(valueJSON))
		if err != nil {
			return err
		}
	}

	return nil
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
	// data, err := a.Read()
	// if err != nil {
	// 	return err, err.Error()
	// }

	// a.Cache = NewCache([][]string{{data}})
	return nil, "stored to the cache!"
}
