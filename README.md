# Go Race Condition Example

This repository demonstrates a common race condition in Go when working with concurrent goroutines and shared variables.

The `bug.go` file contains a program that launches 1000 goroutines, each incrementing a shared counter.  Due to the lack of proper synchronization, the final counter value is often less than 1000, showcasing the race condition.

The `bugSolution.go` file provides a corrected version using a mutex to synchronize access to the counter, ensuring accurate results.

## Running the Code

1. Clone the repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to see the buggy version.
4. Run `go run bugSolution.go` to see the corrected version.