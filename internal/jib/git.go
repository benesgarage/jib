package jib

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
	"time"
)

type IssueIdentifierNotFoundError struct {
	When time.Time
	What string
}

func (e IssueIdentifierNotFoundError) Error() string  {
	return e.What
}

type RemoteNotFoundError struct {
	When time.Time
	What string
}

func (e RemoteNotFoundError) Error() string  {
	return e.What
}


func GetRepositoryFromWD() (*git.Repository, error) {
	wd, err := os.Getwd()
	if nil != err {
		return nil, err
	}

	r, err := git.PlainOpen(wd)
	if nil != err {
		return nil, err
	}

	return r, nil
}
