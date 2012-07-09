# ticked

Timer library.

## Install

```sh
go get github.com/mfine/ticked
```

## Use

```go
package main

import (
	"github.com/mfine/ticked"
	"time"
)

func main() {

	// hello timer every 2 seconds
	helloTimer := ticked.NewTimer(2*time.Second, func() bool { return true }, func() { println("hello") })

	// bye timer every 4 seconds
	byeTimer := ticked.NewTimer(4*time.Second, func() bool { return true }, func() { println("bye") })

	go helloTimer.Run()

	byeTimer.Run()
}
```

