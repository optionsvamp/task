# Task Library

![Build](https://github.com/optionsvamp/task/actions/workflows/build.yaml/badge.svg)

This is a simplistic task-running library in Go. It provides two main components: `Queue` and `Runner`.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
    - [Queue](#queue)
    - [Runner](#runner)
- [Contributing](#contributing)
- [Versioning](#versioning)
- [License](#license)

## Prerequisites
Before you begin, ensure you have met the following requirements:
* You have installed Go version 1.16 or later.

## Installation
To use this library in your Go project, you can import it like any other Go package:

```go
import "github.com/optionsvamp/task"
```

## Usage

This section provides more detailed examples of how to use the Queue and Runner components in your Go project.

### Queue

The `Queue` struct allows you to run multiple tasks concurrently with a specified maximum number of tasks running at the same time.  You
can add to the queue while it's processing its workload.

Here's a basic example of how to use `Queue`:

```go
// Create a new queue that can run up to 2 tasks concurrently
q := NewQueue(2)

// Define a task. This could be any function that matches the Task function signature.
task1 := func(ctx context.Context) {
  // Your task code here. This could be anything from a computation to a network request.
  fmt.Println("Task 1 is running")
}

task2 := func(ctx context.Context) {
  // Another task
  fmt.Println("Task 2 is running")
}

task3 := func(ctx context.Context) {
  // Yet another task
  fmt.Println("Task 3 is running")
  // Simulate a task that takes time by sleeping for 2 seconds
  time.Sleep(2 * time.Second)
}

// Run the tasks without a timeout
q.RunTask(task1, 0)
q.RunTask(task2, 0)

// Run the task with a 1 second timeout to simulate a task that takes too long
q.RunTask(task3, 1*time.Second)

// Wait for all tasks to complete
q.Wait()
```

### Runner

The `Runner` allows you to run multiple tasks sequentially or concurrently with a specified maximum number of tasks running at the same time.

Here's a basic example of how to use Runner:

```go
// Create a new runner
r := NewRunner()

// Define a task. This could be any function that matches the Task function signature.
task1 := func(ctx context.Context) {
// Your task code here. This could be anything from a computation to a network request.
fmt.Println("Task 1 is running")
}

task2 := func(ctx context.Context) {
// Another task
fmt.Println("Task 2 is running")
}

task3 := func(ctx context.Context) {
// Yet another task
fmt.Println("Task 3 is running")
}

// Add the tasks to the runner. You can add as many tasks as you want.
r.AddTask(task1)
r.AddTask(task2)
r.AddTask(task3)

// Run the tasks sequentially without a timeout. This will run the tasks one after the other in the order they were added.
r.RunSequential(0)

// Run the tasks sequentially with a 50 millisecond timeout. If a task takes longer than 50 milliseconds to complete, it will be cancelled.
r.RunSequential(50*time.Millisecond)

// Run the tasks in parallel without a timeout. This will run up to 2 tasks at the same time.
r.RunParallel(2, 0)

// Run the tasks in parallel with a 50 millisecond timeout. This will run up to 2 tasks at the same time, and if a task takes longer than 50 milliseconds to complete, it will be cancelled.
r.RunParallel(2, 50*time.Millisecond)
```

Please note that the RunSequential and RunParallel methods will run the tasks that have been added to the runner using the AddTask method. The RunSequential method will run the tasks one after the other, while the RunParallel method will run the tasks concurrently.  The number of concurrent tasks in RunParallel is determined by the first parameter.  

## Contributing

If you want to contribute to this project, please fork the repository and create a pull request, or open an issue for any bugs or feature requests

## Versioning
We use SemVer for versioning. For the versions available, see the tags on this repository.

## License

This project is licensed under the Unlicense - see the LICENSE.md file for details.