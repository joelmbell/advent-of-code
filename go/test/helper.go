package test

import (
	"errors"
	"fmt"
)

type Case[I any, O comparable] struct {
	Input  I
	Output O
}

func Execute[I any, O comparable](cases []Case[I, O], solver func(I) O) error {
	for i, test := range cases {
		output := solver(test.Input)
		if output != test.Output {
			return errors.New(fmt.Sprintf("test: %v, output: %v != expected: %v", i, output, test.Output))
		}
	}
	return nil
}
