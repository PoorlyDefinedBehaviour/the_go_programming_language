package memo4

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

func httpGetBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http.Get request failed")
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	result result
	ready  chan struct{} // closed when result is ready
}

type Memo struct {
	f     Func
	mutex sync.Mutex
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*result), mutex: sync.Mutex{}}
}

// this is concurrency safe but may affect performance,
// a multiple readers mutex would be better.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mutex.Lock()

	entry := memo.cache[key]

	if entry == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		entry = &entry{ready: make(chan struct{})}

		// NOTE: would all this be necessary if we just initialized entry
		// before putting it in the map?
		memo.cache[key] = entry
		memo.mutex.Unlock()

		entry.result.value, entry.result.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		memo.mutex.Unlock()

		<-e.ready // wait for ready condition
	}

	return entry.result.value, errors.Wrap(entry.result.err, "memo.Get")
}
