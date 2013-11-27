package gotang

import (
	"net/http"
  "encoding/json"
)

func GetJSON(url string, obj interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
