package dao

import "github.com/pkg/errors"

// @generate DO NOT EDIT

func ZGetCriticalUserName() ([]string, error) {

	// step 1
	var success bool
	// process query ...
	// query slave DB
	if success {
		return []string{"Tom", "John"}, nil
	}

	// step 2
	// query master DB
	if success {
		return []string{"Tom", "John"}, nil
	}

	return nil, errors.Wrap(ErrNoRows, "master DB doesn't exist the records.")
}

func ZGetNormalUserName() ([]string, error) {

	// step 1
	var success bool
	// process query ...
	// query slave DB
	if success {
		return []string{"Tom", "John"}, nil
	}

	return nil, errors.Wrap(ErrNoRows, "slave DB doesn't exist the records.")
}
