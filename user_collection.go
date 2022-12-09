package main

import (
	"github.com/martikan/sdk/collections"
)

type userCollection struct {
	users []*user
}

func (u *userCollection) CreateIterator() collections.Iterator[*user] {
	return &userIterator{
		users: u.users,
	}
}
