package service_test

import (
	"context"
	"math/rand"

	"github.com/Jeffail/benthos/v3/public/x/service"
)

type GibberishInput struct {
	length int
}

func (g *GibberishInput) Connect(ctx context.Context) error {
	return nil
}

func (g *GibberishInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	b := make([]byte, g.length)
	for k := range b {
		b[k] = byte((rand.Int() % 94) + 32)
	}
	return service.NewMessage(b), func(ctx context.Context, err error) error {
		// We don't care about acks in this particular input, so do nothing
		return nil
	}, nil
}

func (g *GibberishInput) Close(ctx context.Context) error {
	return nil
}

// This example demonstrates how to create an input plugin, which is configured
// by providing a struct containing the fields to be parsed from within the
// Benthos configuration.
func Example_inputPlugin() {
	configSpec := service.NewConfigSpec().
		Summary("Creates a load of gibberish, putting us all out of work.").
		Field(service.NewIntField("length").Default(100))

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
		length, err := conf.FieldInt("length")
		if err != nil {
			return nil, err
		}
		return &GibberishInput{length}, nil
	}

	err := service.RegisterInput("gibberish", configSpec, constructor)
	if err != nil {
		panic(err)
	}

	// And then execute Benthos with:
	// service.RunCLI()
}
