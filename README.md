# WPS Function Prefix Linter

A Go linter that checks function names start with `WPS_`.

## Installation

``` bash
go install github.com/Ju-DeCo/wpsfuncprefix-linter@latest
```

## Usage

### Standalone

``` bash
wpsfuncprefix ./...
```

### With go vet

``` bash
go vet -vettool=$(which wpsfuncprefix) ./...
```

## Rules

* All functions (except init and main) must start with WPS_
* Test files (_test.go) are ignored

## Example

### Correct

``` go
func WPS_MyFunction() {}
```

### Incorrect
``` go
func MyFunction() {}  // will trigger lint error
```