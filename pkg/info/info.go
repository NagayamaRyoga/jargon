package info

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

func Encode[T any](v *T) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func Decode[T any](s string) (*T, error) {
	if len(s) == 0 {
		return nil, nil
	}

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var v T
	dec := gob.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}

	return &v, nil
}
