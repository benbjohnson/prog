Go: testing
===========

Testing in Go is very limited compared to some other languages, however,
the limited support is intentional. Rather than building a complex testing
language or a full featured testing library, Go uses the Go language itself.

The `testing` package in Go is broken down into 3 parts:

1. Tests

2. Benchmarks

3. Examples


## Tests

Tests are simple functions that begin with the word `Test` and accept a
`*testing.T` object that is used to control the test. There are a few
operations you can perform with the test object:

* Fail - Marks a test as failed and stops immediately.

* Error - Marks a test as failed but continues the test.

* Skip - Marks a test as skipped and stops immediately.

* Log - Print a message to the test's output buffer.

* Parallel - Indicates that a test can be run in parallel with other tests.

Each of these have helper functions to format log messages but this is
essentially all you can do with a test. The control flow of the test is
handled by the Go language rather than assert statements or test-specific DSLs.


## Benchmarks

Benchmarks are simple functions that begin with the word `Benchmark` and
accept a `*testing.B` object that is used to control the benchmark. A
benchmark largely supports the same operations as a test but also adds a
few additional operations:

* N - This is the number of iterations to run in the benchmark.

* ReportAllocs() - Enable the memory allocation reporting for the benchmark.

* SetBytes() - Enable IO reporting for the benchmark.

* ResetTimer() - Restarts the start time of the benchmark. Useful for setup code.

The Go toolchain will automatically adjust the number of iterations, `N`,
until it can run a benchmark function for a minimum duration.


## Examples

Examples are functions that demonstrate some functionality of your packages
that will show up in godoc. The great part about examples is that Go will
run and test your example so you know that godoc examples should always work!



