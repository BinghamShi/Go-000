package main

import (
	"errors"

	xerrors "github.com/pkg/errors"
)

func processUser() error {
	_, err := dao.GetUserName()
	if err != nil {
		if errors.Is(err, dao.ErrNoRows) {
			// process no records
			return err
		}
		// the process of other DB issue
		return err
	}
	return nil
}

func processZUser() error {
	_, err := dao.GetZUserName()
	if err != nil {
		return xerrors.Wrap(err, "get user name failed")
	}
	return nil
}

func processSpecialUser(isVIP bool) error {

	if isVIP {
		_, err := dao.ZGetCriticalUserName()
		if err != nil {
			if errors.Is(err, dao.ErrNoRows) {
				// process no records
				return err
			}
			// the process of other DB issue
			return err
		}
		return nil
	}

	_, err := dao.ZGetNormalUserName()
	if err != nil {
		if errors.Is(err, dao.ErrNoRows) {
			// process no records
			return err
		}
		// the process of other DB issue
		return err
	}
	return nil

}
