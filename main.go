package main

import "fmt"

func main() {
	// Sample for non-blocking channels.
	nonBlockingChannel()

	// Sample to use collections.

	user1 := &user{
		id:    1,
		email: "email@example.com",
	}

	user2 := &user{
		id:    2,
		email: "email2@example.com",
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User email is %s\n", user.email)
	}

}
