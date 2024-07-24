package main

import (
	"fmt"

	// errs "github.com/pkg/errors"
	"github.com/andrewbulin/errexample/thework"
)

func main() {
	for i := 0; i < 10; i++ {

		err := thework.Stuff{}.Do()

		re := handleErr(err)

		switch {
		case re == nil:
			fmt.Printf("result %d: okay!\n", i)
		default:
			fmt.Printf("result %d: %v\n", i, re)
		}
	}
}

// NOTE: logic at this layer of the app to handle the errors
// in the right way for the current context

// evaluator is an interface that defines a method to evaluate an error,
// which is implemented by the internal `evalerr` package,
// without the need to expose or import it.
type evaluator interface {
	Evaluate() error
}

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	e, ok := err.(evaluator)
	if !ok {
		// do stuff to handle an error that we can't evaluate

		return err
	}

	ee := e.Evaluate()

	// do stuff to handle an error that we can evalute.
	// for example, we could take just the root of the error,
	// and make decisions based on it
	// root := errs.Cause(ee)

	// for this example, let's just return the evaluted error
	return ee
}
