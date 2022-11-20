package cloudfunc

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)

func HttpError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		log.Println(err)
	}
	http.Error(w, err.Error(), status)
	w.Write([]byte(fmt.Sprintf("REQUEST FAILED: %d %s", status, err)))
}

func ParseJSON(r *http.Request, dst interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		r.Body.Close()
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}

func ServeJSON(w http.ResponseWriter, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}