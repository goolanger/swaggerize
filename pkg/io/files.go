package io

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"os"
)

func Save(specs *swagger.Instance, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	encoded, err := Encode(specs)
	if err != nil {
		return err
	}

	_, err = file.Write(encoded)
	if err != nil {
		return err
	}

	return nil
}
