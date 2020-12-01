package dao

import "github.com/pkg/errors"

func GetUserName() ([]string, error) {
	var useless bool

	if useless {
		return nil, errors.Wrap(ErrNoRows, "useless return")
	}

	return nil, errors.Wrap(ErrNoRows, "normal return")
}
