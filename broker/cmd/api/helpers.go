package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxByteSize := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByteSize))

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have one single value")
	}

	return nil
}
