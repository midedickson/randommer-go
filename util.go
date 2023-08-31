package randommer

import (
	"encoding/json"
	"io"
	"log"
)

func parseRequestBody(BodyStream io.Reader, target interface{}) error {
	body, _ := io.ReadAll(BodyStream)
	log.Println(string(body))
	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return nil

}
