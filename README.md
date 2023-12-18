# Checkout test exercise

## Background

Testing is crucial at Ascenda. We cover all of our code in unit tests,
and we make sure to provide integration tests for critical parts of our systems.

This exercise presents a very simplified checkout process. It consists of 3 steps:

1. Validate parameters
  - If validation fails, return an error
2. Perform booking with external supplier
  - When booking fails, return an error
3. Create booking object

You are provided 2 structs that you should use:

1. `checkout` performs the steps above and returns a booking or raises error
2. `application` allows to create and retrieve bookings

## Task

Your task is to write a set of unit tests for the attached `app.go` and `checkout.go` file.

Your tests should cover `application` and `checkout`. The `api` is provided only for reference.

You are provided a `tests/app/app_test.go` and `tests/app/checkout_test.go` file that you should modify. You can use another testing file structures or packages if you prefer.

After you deliver your code, I'll try to break your tests in a following way:

* I will be modifying application code to make it stop working correctly
* After each change I'll run the tests to see if your tests detected that I broke the application
* If your tests still pass, it means that they did not fully cover the application
* If your tests do not pass, it means they covered the part of app that I broke
* The more times your tests fail when I break the app, the better

What matters in the output:

* correctness - your tests should pass when app works, fail when it's broken
* readability - your tests should be easy to understand
* design - the tests must be well separated, remember they are unit tests, which means one test should verify a behaviour of a single class or method

## Delivery

Once you finish the task, send me your updated `app_test.go` and `checkout_test.go` file by email or Github Gist

## Questions?

If you have any questions, drop me an e-mail!
