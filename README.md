# BID - Binary Unique ID

## pgbid - PostgreSQL extension

See to [pgbid](./pgbid) directory.

## Installation

    go get github.com/moisespsena-go/bid

## Sample Usage

```go
package main

import (
	"fmt"

	"github.com/moisespsena-go/bid"
)

type User struct {
	ID bid.BID
	Name string
}

func main() {
	// simple generation
	fmt.Println(bid.New())
	
	// generate from field
	var u User
	u.ID.Generate()
	fmt.Println(u.ID)
}
```

See tests for details.

## LICENSE

MIT License
