package test

import (
	"errors"
	"fmt"
)

type testableOutput interface {
	int64 | int | string
}

type Case[I any, O testableOutput] struct {
	Input  I
	Output O
}

func Execute[I any, O testableOutput](cases []Case[I, O], solver func(I) O) error {
	for i, test := range cases {
		output := solver(test.Input)
		if output != test.Output {
			return errors.New(fmt.Sprintf("test: %v, output: %v != expected: %v", i, output, test.Output))
		}
	}
	return nil
}
