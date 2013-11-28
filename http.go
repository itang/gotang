package gotang

import (
	"encoding/json"
	"net/http"
)

func GetJSON(url string, obj interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}
