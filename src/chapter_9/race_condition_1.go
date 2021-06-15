package main

import "image"

// A race condition occurs whenever two or more goroutines access access
// the same variable concurrently and at least one of the accesses is a write.

// ---
// race condition example
//
var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	// load icon from somewhere
	panic("make compiler happy")
}

// This function is not concurrency-safe because
// there is a data race accessing the map.
func Icon(name string) image.Image {
	if _, ok := icons[name]; !ok {
		// lazy loading icons
		icons[name] = loadIcon(name)
	}

	return icons[name]
}

// ---
// solution
//
// variable is initialized before main() is executed.
var icons = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// This function is concurrency safe because it is read only.
func Icon(name string) image.Image {
	return icons[name]
}

// ---

func main() {

}
