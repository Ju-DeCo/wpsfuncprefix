# WPS Function Prefix Linter

A Go linter that checks function names start with `WPS_`.

## Installation

``` bash
go install github.com/Ju-DeCo/wpsfuncprefix@latest
```

## Usage

### Standalone

``` bash
wpsfuncprefix ./...
```

### With go vet

``` bash
go vet -vettool="$env:yourinstall\path\wpsfuncprefix.exe" ./...
```

### golangci-lint

```yaml
version: "2"

linters:
  default: none
  enable:
    - wpsfuncprefix

  settings:
    custom:
      wpsfuncprefix:
        type: module
        original-url: github.com/Ju-DeCo/wpsfuncprefix
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