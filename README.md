# GO-projects

## The go mod init command creates a go.mod file to track your code's dependencies. 

## example :  go mod init example.com/greetings

### In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type)

## Edit the example.com/hello module to use your local example.com/greetings module

`$ go mod edit -replace example.com/greetings=../greetings`
 
 ### run the go mod tidy command to synchronize the

 `go run .` to run the main program

 ## Error handling

 `if name == "" {
		return "", errors.New("empty name")
	}`

`if err != nil {
		log.Fatal(err)
	}`

## In Go, you initialize a map with the following syntax: make(map[key-type]value-type).

## Ending a file's name with _test.go tells the go test command that this file contains test functions.

## The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.

##  run the `go build` command to compile the code into an executable.

##### You've compiled the application into an executable so you can run it. But to run it currently, your prompt needs either to be in the executable's directory, or to specify the executable's path.

## You can discover the install path by running the go list command, as in the following example:

`$ go list -f '{{.Target}}'`

## Once you've updated the shell path, run the go install command to compile and install the package.

`$ go install`



