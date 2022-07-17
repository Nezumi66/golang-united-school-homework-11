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
		func(i int64) {
			errGroup.Go(func() error {
				user := getOne(i)
				mu.Lock()
				result = append(result, user)
				mu.Unlock()
				return nil
			})
		}(i)
	}

	errGroup.Wait()

	return result
}
