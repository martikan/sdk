package main

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) GetNext() *user {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++

		return user
	}

	return nil
}

func (u *userIterator) GetAllElements() []*user {
	return u.users
}
