package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var users []user
	var waitGroup sync.WaitGroup
	var i int64

	for i = 0; i < n; i++ {
		waitGroup.Add(1)

		go func(i int64) {
			u := getOne(i)
			users = append(users, u)
			waitGroup.Done()
		}(i)
	}

	return users
}
