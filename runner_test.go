package task

import (
	"context"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	r := NewRunner()

	// Define two tasks that sleep for different durations
	task1 := func(ctx context.Context) error {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-ctx.Done():
		}
		return nil
	}
	task2 := func(ctx context.Context) error {
		select {
		case <-time.After(200 * time.Millisecond):
		case <-ctx.Done():
		}
		return nil
	}

	// Add the tasks to the runner
	r.AddTask(task1)
	r.AddTask(task2)

	// Run the tasks sequentially without a timeout
	start := time.Now()
	err := r.RunSequential(0)
	duration := time.Since(start)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if duration < 300*time.Millisecond {
		t.Errorf("Expected tasks to run for at least 300ms, but they ran for %v", duration)
	}

	// Run the tasks in parallel without a timeout
	start = time.Now()
	err = r.RunParallel(2, 0)
	duration = time.Since(start)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if duration > 210*time.Millisecond {
		t.Errorf("Expected tasks to run for no more than 210ms, but they ran for %v", duration)
	}
}
