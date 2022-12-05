# advent_of_code_2022

My solutions (or attempts) to Advent of Code 2022 using Go!

## Setup
- Clone the repository
- Install Go:
  - For mac: `brew install go`
  - For windows: ...go to https://go.dev I suppose?

  
## Structure
You'll find a `main.go`, this contains the main function to initiate the running of tasks and some error handling/validation on user input. 
This file defines an interface for days to adhear to.

Inside the `days/` package you'll find individual packages for each day e.g. `days/day1/`. These contain a main file for the day's tasks and a
corresponding test file.

There is a `utils/` package too, this holds some helper functions (mostly just reading an input file and converting to [][]string)

Finally there is an `inputs/` package, here the .txt file inputs are found. The file naming is opinionated, the given input is `day{dayNumber}.txt`
and then any corresponding tests will be `day{dayNumber}_test{testNumber}.txt`

## Running
#### Run from terminal in project root

- To simply run the program for a given day
  - `go run . {chosen day}`

- To run the tests
  - For all tests: `go test ./... -v`
  - For specific day: `go test ./day{chosen day} -v`
  
- Or, if you feel like it, you could build the project to get an executable
  - `go build`
  - Then you can run the program file with `./advent_of_code_2022 {chosen day}`
     _note: you will need the input .txt files in the right location relative to the .exe, `/inputs/day{number}/`_

### Have fun!

