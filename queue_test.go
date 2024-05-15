package task

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue(2)

	// Define a simple task that sleeps for a specified duration
	task := func(ctx context.Context) error {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-ctx.Done():
		}
		return nil
	}

	// Run the task without a timeout
	start := time.Now()
	q.RunTask(task, 0)
	q.RunTask(task, 0)
	q.Wait()
	duration := time.Since(start)
	if duration < 100*time.Millisecond {
		t.Errorf("Expected tasks to run for at least 100ms, but they ran for %v", duration)
	}
	if duration > 210*time.Millisecond {
		t.Errorf("Expected tasks to run for no more than 210ms, but they ran for %v", duration)
	}

	// Run the task with a timeout
	start = time.Now()
	q.RunTask(task, 50*time.Millisecond)
	q.RunTask(task, 50*time.Millisecond)
	q.Wait()
	duration = time.Since(start)
	if duration > 110*time.Millisecond {
		t.Errorf("Expected tasks to run for no more than 110ms, but they ran for %v", duration)
	}

	// Run the task with a timeout that will be exceeded
	start = time.Now()
	q.RunTask(task, 10*time.Millisecond)
	q.RunTask(task, 10*time.Millisecond)
	q.Wait()
	duration = time.Since(start)
	if duration < 10*time.Millisecond {
		t.Errorf("Expected tasks to run for at least 10ms, but they ran for %v", duration)
	}
	if duration > 30*time.Millisecond {
		t.Errorf("Expected tasks to run for no more than 30ms, but they ran for %v", duration)
	}
}

func TestQueue_ErrorTask(t *testing.T) {
	q := NewQueue(2)

	// Define a task that returns an error
	errorTask := func(ctx context.Context) error {
		return errors.New("task error")
	}

	// Run the error task
	err := q.RunTask(errorTask, 0)
	if err == nil || err.Error() != "task error" {
		t.Errorf("Expected error 'task error', but got %v", err)
	}
}
