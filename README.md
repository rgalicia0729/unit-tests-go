# Unit Tests Go
Example of unit tests in go

# Test commands

- Run test

    $ go test

- Test a specific directory

    $ go test -v ./directory

- Test coverage

    $ go test -cover -v ./directory

    $ go test -coverprofile=coverage.out -v ./directory

    $ go tool cover -func=coverage.out

    $ go tool cover -html=coverage.out

    $ go test -cpuprofile=profile.out -v ./directory

    $ go tool pprof profile.out (top, list nameFunction, web)
