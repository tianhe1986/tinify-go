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
	
	// from file
	tinify.FromFile("./1.jpg").ToFile("./test1.jpg")
	
	// from url
	tinify.FromUrl("https://XXXXXX.png").ToFile("./test2.png")
}
```

## Resize

There are 4 ways your image will be resized according to the [official document](https://tinypng.com/developers/reference/php)

So you could call `Resize` with different method name, such as follows:
```
package main

import (
	"fmt"
	"github.com/tianhe1986/tinify-go/tinify"
)

func main() {
	var key string = "XXXXXX"
	tinify.SetKey(key)
	
	// scale with width
	tinify.FromFile("./1.jpg").Resize("scale", 200, 0).ToFile("./test1-scale-width.jpg")

	// scale with height
	tinify.FromFile("./1.jpg").Resize("scale", 0, 400).ToFile("./test1-scale-height.jpg")

	// fit
	tinify.FromFile("./1.jpg").Resize("fit", 200, 200).ToFile("./test1-fit.jpg")

	// cover
	tinify.FromFile("./1.jpg").Resize("cover", 300, 300).ToFile("./test1-cover.jpg")

	// thumb
	tinify.FromFile("./1.jpg").Resize("thumb", 250, 250).ToFile("./test1-thumb.jpg")
}
```