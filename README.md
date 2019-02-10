**This is simple package for ratesapi.io**

**Usage**
```go
package main

import (
    "fmt"
    "github.com/zeroabe/ratesapi-go"
)


func main() {
    base := "USD"
    symbA, symbB := "EUR", "GPP"
    ratesResponse, err := ratesapi_go.LatestByBaseAndSymbols(base, symbA, symbB)
    if err != nil {
        //handle it
    }
    fmt.Printf("rates:\n %+v", ratesResponse)
}
```

You can use as much symbols as you wish.

***More methods later, folks!***