package steam

import (
	"encoding/json"
	"io"
)

func ParseJson(b io.Reader, v interface{}) error {
	bb, err := io.ReadAll(b)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bb, v)
	if err != nil {
		return err
	}

	return nil
}
