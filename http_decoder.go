package utils

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
)

func JsonDecode(resp io.ReadCloser, model any) error {
	if reflect.ValueOf(model).Kind() != reflect.Pointer {
		return errors.New("value provided is not a pointer")
	}

	data, err := io.ReadAll(resp)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, model); err != nil {
		return err
	}

	return nil
}
