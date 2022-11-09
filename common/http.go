package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func (app *App) GetJSON(url string, dst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if dst != nil {
		if err := json.Unmarshal(b, dst); err != nil {
			return err
		}
	}
	return nil
}
