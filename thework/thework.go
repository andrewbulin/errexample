package thework

import (
	"errors"
	"math/rand"

	"github.com/andrewbulin/errexample/evalerr"
)

type Stuff struct{}

// Do does stuff, for now it invokes a randomizer to determine if it should
// return an error, and if so, select one from the pkg evalerr.
// Pick a random number, 0-2 will return an evalerr, 3 is is an unexpected
// error, and everything else is no error. This is just for the sake of
// the example, in a real world scenario, this would be a more complex logic.
func (s Stuff) Do() error {
	n := rand.Intn(10)
	switch n {
	case 0:
		return evalerr.YellowCardError{}

	case 1:
		return evalerr.RedCardError{}

	case 2:
		return evalerr.BlackCardError{}

	case 3:
		return errors.New("beep beep boop - unexpected error")

	default:
		return nil
	}
}
