package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	e := json.Unmarshal([]byte(body), x)
	if e != nil {
		log.Fatal(e)
		return
	}
}
