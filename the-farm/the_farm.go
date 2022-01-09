package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	count int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", s.count)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fooder, err := weightFodder.FodderAmount()
	if err != nil {
		if fooder >= 0 {
			if err == ErrScaleMalfunction {
				return (fooder * 2.0) / float64(cows), nil
			}
		} else if fooder < 0 {
			if err == ErrScaleMalfunction {
				return 0.0, errors.New("negative fodder")
			}
			return 0.0, errors.New("non-scale error")
		}
		return 0.0, err
	}
	if fooder < 0 {
		return 0.0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{
			count: cows,
		}
	} else {
		perCow := fooder / float64(cows)
		return perCow, nil
	}
}
