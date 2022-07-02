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
	sem := make(chan struct{}, pool)
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(int(n))
	var i int64
	for i = 0; i < n; i++ {
		sem <- struct{}{}
		go func(i int64) {
			user := getOne(i)
			mutex.Lock()
			res = append(res, user)
			mutex.Unlock()
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
