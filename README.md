# env

The env package provides a generic way to retrieve environment variables in Go. It includes two functions, GetEnv and GetEnvList, that can be used to retrieve environment variables of any data type or list of data types respectively. If the specified environment variable is not found, a default value is returned. The package also includes testing functions, TestGetEnv and TestGetEnvList, that can be used to test the conversion of environment variable values to the specified data type.

## Installation

```bash
go get -u github.com/ginger-go/env
```

## Documentation

Please refer to [https://ginger-go.gitbook.io/env/](https://ginger-go.gitbook.io/env/) for the documentation.

## Perform test

Run ginkgo tests with the following command:
```bash
ginkgo -r -v -p --cover --coverpkg=github.com/ginger-go/env
```

Read the coverage report with the following command:
```bash
go tool cover -html=coverprofile.out
```