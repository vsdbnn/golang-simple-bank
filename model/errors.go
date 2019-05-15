package model

import (
	"errors"
)

var (
	ErrAlreadyExists    = errors.New("account already exists")
	ErrNotFound         = errors.New("account not found")
	ErrSelfPayment      = errors.New("self payment")
	ErrLowBalance       = errors.New("not enough money")
	ErrIncorrectBalance = errors.New("incorrect balance")
	ErrNegativePayment  = errors.New("money amount is negative or zero")
)
