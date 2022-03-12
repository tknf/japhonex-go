# Japhonex Go

Japanese phone number checker for Go.

## Installation
```bash
$ go get github.com/tknf/japhonex-go
```

## Usage
In your Go app you can do something like:
```go
import (
  "fmt"

  "github.com/tknf/japhonex-go"
)

func main() {
  result, err := japhonex.OptionalHyphen("<Phone number input>")

  if err != nil {
    fmt.Println(err)
  }

  if result == true {
    // validation is passed
  }
}
```

### Hyphen validation patterns
### Optional
```go
result, err := japhonex.OptionalHyphen("<Input>")
// 0xx-xxxx-xxxx or 0xxxxxxxxxx
```
### Required
```go
result, err := japhonex.RequireHyphen("<Input>")
// 0xx-xxxx-xxxx
```
### No hyphen
```go
result, err := japhonex.NoHyphen("<Input>")
// 0xxxxxxxxxx
```

## Licence
MIT