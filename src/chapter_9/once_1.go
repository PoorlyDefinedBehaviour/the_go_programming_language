package once1

import (
	"image"
	"sync"
)

var icons map[string]image.Image

func loadIcon(name string) image.Image {
	// pretend to load icon from somewhere
	panic("just to make the compiler happy")
}

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// This function is not concurrency-safe because in the
// absence of explicit synchronization, the compiler and CPU
// are free to reorder accesses to memory in any number of ways,
// so as long the behaviour of each goroutine is sequentially consistent.
//
// Reordering example that makes this function not behave as expected:
//
// func loadIcons() {
// 	 icons = make(map[string]image.Image)
// 	 icons["spades.png"] = loadIcon("spades.png")
// 	 icons["hearts.png"] = loadIcon("hearts.png")
// 	 icons["diamonds.png"] = loadIcon("diamonds.png")
// 	 icons["clubs.png"] = loadIcon("clubs.png")
// }
//
// Given this order, a goroutine may enter concurrencyUnsafeIcon()
// and find icons to be not nil and try to accesses one of it's keys,
// but icons being not nil does not guarantee that all keys are initialized.
func concurrencyUnsafeIcon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}

	return icons[name]
}

var withMutexIconMutex sync.Mutex

// Concurrency-safe because of the mutex.
func WithMutexIcon(name string) image.Image {
	withMutexIconMutex.Lock()
	defer withMutexIconMutex.Unlock()

	if icons == nil {
		loadIcons()
	}

	return icons[name]
}

// This is concurrency-safe but too ugly and error prone.
var withRWMutexIconMutex sync.RWMutex

func WithRWMutexIcon(name string) image.Image {
	// acuire read lock
	withRWMutexIconMutex.RLock()

	if icons != nil {
		icon := icons[name]
		withRWMutexIconMutex.RUnlock()
		return icon
	}

	withRWMutexIconMutex.RUnlock()

	// acquire write lock
	withRWMutexIconMutex.Lock()
	defer withRWMutexIconMutex.Unlock()

	if icons == nil {
		loadIcons()
	}

	return icons[name]
}

// This is concurrency-safe.
var loadIconsOnce sync.Once

func WithOnceIcon(name string) image.Image {
	// sync.Once is a struct that holds a mutex and flag
	// to check wether initialization has happened.
	// The function given to sync.Once.Do(f) is only called once.
	// If multiple goroutines call sync.Once.Do(f) at the same time,
	// all goroutines wait until f() is done executing but f() is only called once.
	// Subsequent calls have no effect.
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
