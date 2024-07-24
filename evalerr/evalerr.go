package evalerr

import (
	"errors"

	errs "github.com/pkg/errors"
)

type EvalError int

const (
	YellowCard EvalError = iota
	RedCard
	BlackCard
)

var (
	ErrYellowCardSituation = errors.New("yellow card situation")
	ErrRedCardSituation    = errors.New("red card situation")
	ErrBlackCardSituation  = errors.New("black card situation")
)

type YellowCardError struct {
	// details
}

func (e YellowCardError) Error() string {
	return ErrYellowCardSituation.Error()
}

func (e YellowCardError) Evaluate() error {
	// doing stuff with the error details
	msg := "This is a " + e.Error() + "! Be careful..."
	err := errs.Wrap(ErrYellowCardSituation, msg)

	return err
}

type RedCardError struct {
	// details
}

func (e RedCardError) Error() string {
	return ErrRedCardSituation.Error()
}

func (e RedCardError) Evaluate() error {
	// doing stuff with the error details
	msg := "This is a " + e.Error() + "! You're out!"
	err := errs.Wrap(ErrRedCardSituation, msg)

	return err
}

type BlackCardError struct {
	// details
}

func (e BlackCardError) Error() string {
	return ErrBlackCardSituation.Error()
}

func (e BlackCardError) Evaluate() error {
	// doing stuff with the error details
	msg := "This is a " + e.Error() + "! You never learn, do you?!"
	err := errs.Wrap(ErrBlackCardSituation, msg)

	return err
}
