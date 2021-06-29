package main

// The go build tool treats a package specially if its import
// path contains a path segment named internal.
// Such packages are called internal packages. An internal
// package may be imported only by another package that is inside
// the tree rooted at the parent of the internal directory.
// For example, given the packages below, net/http/internal/chunked
// can be imported from net/http/httputil or net/http, but from
// net/url. However, net/url may import net/http/httputil.
//
// net/http
// net/http/internal/chunked
// net/http/httputil
// net/url

func main() {}
