package io

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"os"
)

func Save(specs *swagger.Instance, path string) error {


	f, err := os.Open(path)
	if err != nil {
		f, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	rep, err := Encode(specs)
	if err != nil {
		return err
	}

	_, err = f.Write(rep)
	if err != nil {
		return err
	}

	return nil
}
