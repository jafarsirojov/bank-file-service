package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ReadJSONBody(request *http.Request, dto interface{}) (err error) {
	if request.Header.Get("Content-Type") != "application/json" {
		return errors.New("error")
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New("error")
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return errors.New("error")
	}
	return nil
}
