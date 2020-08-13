package test

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/pkg/io"
	"testing"
)

func InitTest(t*testing.T) {
	specs := swagger.New()

	err := io.Save(specs)

	if err==nil {
		t.Fatal(err)
	}
}
