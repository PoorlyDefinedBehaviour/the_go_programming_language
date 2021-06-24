package memo1

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

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// this is not concurrency-safe
func (memo *Memo) Get(key string) (interface{}, error) {
	result, ok := memo.cache[key]
	if !ok {
		result.value, result.err = memo.f(key)
		memo.cache[key] = result
	}

	return result.value, result.err
}
