
# Overview

# Unit Tests

For the unit tests, I kept to the Go standard of embedding the test files with the code under test.  In this case, I expanded the already present calculator_test.go file.  I did rearrange the tests such that each  operator was tested independently.  This made it clear what tests were being done on each operator.

A number of the tests are commented out, due to not properly handling overflows and the usual issues with doing floating point math.

# Integration Tests

The integration tests are located in the /integ-tests directory.  The unit tests extensively test the correctness of the various operations, so the integration tests focus on the output.  These tests are dependent on the calculator being built and capture stdout and stderr for validation.  They also use testify package which provides a set of assert functions that make testing cleaner.

# Running the tests

The tests can be run using make.  make unit-test to run the unit tests; make integ-test to run the integration tests and make test will run all tests.

## CI/CD

The CI/CD for this project uses github actions.  When a PR is created on branch off of main, the action runs automatically.  If it fails, the PR can’t be merged.
