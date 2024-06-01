# FizzBuzz Service

This service provides two main functionalities:

## SingleFizzBuzz Function

This function behaves as follows:

1. It accepts an integer number `n` and by default, it returns the integer number `n` without any operation.
2. If `n` is divisible by 3, it returns `Fizz`.
3. If `n` is divisible by 5, it returns `Buzz`.
4. If `n` is divisible by both 3 and 5, it returns `FizzBuzz`.

The implementation can be found in `fizzbuzz/fizzbuzz.go`.

## GET /range-fizzbuzz Endpoint

This HTTP endpoint has the following requirements:

1. It accepts two parameters: `from` and `to`. Both should be integers, and `from` should be less than or equal to `to`.
2. The response returns the value of `SingleFizzBuzz` for each integer between `from` and `to` (inclusive), delimited by a space.

The implementation of the `RangeFizzBuzz` function can be found in `fizzbuzz/fizzbuzz.go`, and the handler is in `main.go`, the server run on port 8080.

## Additional Features

The service also includes the following features:

1. It logs its request, response, and latency to STDOUT.
2. It has performance requirements for the endpoint:
   - 1 second as timeout
   - Can create at maximum 1000 goroutines for the calculation at the same time for all requests
   - Accepts at maximum 100 numbers as the range
3. It can be terminated gracefully using SIGTERM.

The logging, timeout, and maximum 1000 goroutine for request are implemented as middleware in `middleware/middleware.go`.

# Execute
To execute the app, run: `make run`

# Test
To run test, execute `make test`