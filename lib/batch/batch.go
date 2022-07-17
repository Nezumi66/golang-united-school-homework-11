package batch

import (
	"golang.org/x/sync/errgroup"
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
	errGroup := errgroup.Group{}
	errGroup.SetLimit(int(pool))

	var i int64
	var result []user
	mu := sync.Mutex{}

	for i = 0; i < n; i++ {
		go func() {

		}()
	}

	errGroup.Wait()

	return result
}
