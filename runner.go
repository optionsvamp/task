package task

import (
	"context"
	"sync"
	"time"
)

type Runner struct {
	tasks []Task
}

func NewRunner() *Runner {
	return &Runner{}
}

func (tr *Runner) AddTask(task Task) {
	tr.tasks = append(tr.tasks, task)
}

func (tr *Runner) RunSequential(timeout time.Duration) error {
	for _, task := range tr.tasks {
		ctx := context.Background()
		if timeout != 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}

		done := make(chan bool, 1)
		var taskErr error
		go func() {
			taskErr = task(ctx)
			done <- true
		}()

		select {
		case <-done:
			// Task completed within the timeout.
		case <-ctx.Done():
			// Task didn't complete within the timeout.
			// Handle the timeout here as needed.
		}

		if taskErr != nil {
			return taskErr
		}
	}
	return nil
}

func (tr *Runner) RunParallel(n int, timeout time.Duration) error {
	var wg sync.WaitGroup
	sem := make(chan bool, n)

	// Shared error variable and its mutex
	var err error
	var errMutex sync.Mutex

	for _, task := range tr.tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			sem <- true

			ctx := context.Background()
			if timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, timeout)
				defer cancel()
			}

			taskErr := task(ctx) // Execute the task and get the error (or nil if no error)

			// If the task returned an error and no other task has set the error yet, store it
			if taskErr != nil {
				errMutex.Lock()
				if err == nil {
					err = taskErr
				}
				errMutex.Unlock()
			}

			<-sem
		}(task)
	}

	wg.Wait()

	return err
}
