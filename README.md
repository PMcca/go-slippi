# go-slippi

`go-slippi` is a Golang (1.21) parser for `.slp` slippi files. It (currently only) supports parsing .slp files into Go 
types.

## Usage

### Full game
```go
package main

import (
	"fmt"
	goslippi "github.com/pmcca/go-slippi"
	"log"
)

func main() {
	filePath := "path/to/my-replay.slp"
	game, err := goslippi.ParseGame(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(game.Data.GameStart.SlippiVersion)
}
```

### Metadata Only
```go
package main

import (
	"fmt"
	goslippi "github.com/pmcca/go-slippi"
	"log"
)

func main() {
	filePath := "path/to/my-replay.slp"
	meta, err := goslippi.ParseMeta(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(meta.Players)
}
```
