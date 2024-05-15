package task

import (
	"context"
	"sync"
	"time"
)

type Queue struct {
	sem chan bool
	wg  *sync.WaitGroup
}

func NewQueue(n int) *Queue {
	return &Queue{
		sem: make(chan bool, n),
		wg:  new(sync.WaitGroup),
	}
}

func (wq *Queue) RunTask(task Task, timeout time.Duration) error {
	wq.wg.Add(1)
	errChan := make(chan error, 1)
	go func(task Task) {
		defer wq.wg.Done()
		wq.sem <- true

		ctx := context.Background()
		if timeout != 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}

		errChan <- task(ctx)

		<-wq.sem
	}(task)

	// Check if the task returned an error
	err := <-errChan
	if err != nil {
		return err
	}

	return nil
}

func (wq *Queue) Wait() {
	wq.wg.Wait()
}
