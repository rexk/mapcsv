# MAPCSV
Simple encoding/csv wrapper that returns map[string]string as a result rather than []string

## Install
```bash
go get install github.com/rexk/mapcsv
```

## Usuage
```go
package main

import (
  "gihtub.com/rexk/mapcsv"
  "encoding/csv"
  "os"
  "fmt"
)

func main() {
  file, err := os.Open("file.csv")
  if err != nil {
    panic(err)
  }

  r := mapcsv.NewReader(csv.NewReader(file))
  records, err := r.ReadAll()
  if err != nil {
    panic(err)
  }

  fmt.Println(records)
}
```
