# Go MRSS

Helper package to generate MRSS media feeds. Built and tested for Live Feeds in Brightsign players

Based on the [Media RSS specification](http://search.yahoo.com/mrss/), but only to the extent that is needed for Live Feeds.

## Installation

To install the package:

```sh
go get -u github.com/carrier-labs/go-mrss
```
Import it into your code:

```go
import "github.com/carrier-labs/go-mrss"
```

## Quick start


```go
package main

import (
  "net/http"

  "github.com/carrier-labs/go-mrss"
)

func main() {
    // Create a new MRSS Feed Object
    mrss := mrss.New("A Title","http://example.com/")
  
    // Add a new item to the feed
    mrss.AddItem(Item{
        Title: "Item Title",
        ...
    })

    fmt.Printf("%s\n", mrss.toString())
}
```
