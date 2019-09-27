# tinify-go
Golang client for the Tinify API

# Usage
## Basic Compressing images

This very simple. Just set your API key, then call with your input filename and output filename. As follows:
```
package main

import (
	"fmt"
	"github.com/tianhe1986/tinify-go/tinify"
)

func main() {
	var key string = "XXXXXX"
	tinify.SetKey(key)
	tinify.FromFile("./1.jpg").ToFile("./test1.jpg")
}
```