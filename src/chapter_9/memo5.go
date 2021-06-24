package memo5

import (
	"io/ioutil"
	"net/http"

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

func (e *entry) call(f Func, key string) {
	e.result.value, e.result.err = f(key)

	// mark as ready
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait until entry is ready.
	<-e.ready

	// Send result to client.
	response <- e.result
}

type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key: key, response: response}
	result := <-response
	return result.value, errors.Wrap(result.err, "memo.Get")
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for request := range memo.requests {
		e := cache[request.key]

		// if this is the first request for this key.
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[request.key] = e
			go e.call(f, request.key)
		}
		go e.deliver(request.response)
	}
}
