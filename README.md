# mysql-kv-wrapper

## Installation

```
go get github.com/DCsunset/mysql-kv-wrapper
```

## Examples

```go
package main

import (
	kvstore "github.com/DCsunset/mysql-kv-wrapper"
)

func main() {
	var store kvstore.KVStore
	err := store.Open("root:password@/")
	if err != nil {
		panic(err)
	}
	defer store.Close()

	err = store.Write("hello", "world")
	if err != nil {
		panic(err)
	}
	value, err := store.Read("hello")
	if value != "world" {
		panic("Wrong value")
	}
}
```
