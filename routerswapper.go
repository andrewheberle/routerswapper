// Package routerswapper implements a simple process to swap a Golang HTTP
// router (gorilla/mux etc) during runtime. This is aimed at allowing
// route/handler changes based on a configuration change or update.
package routerswapper

import (
	"net/http"
	"sync"
)

// router interface is satisfied by any type that implements ServeHTTP
type router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// RouterSwapper is the type used for swapping
type RouterSwapper struct {
	mu sync.RWMutex
	rt router
}

// Swap replaces the current router.
func (rs *RouterSwapper) Swap(rt router) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// New creates a new RouteSwapper based on the provided router
// which can then be used where ServeHTTP would be used
func New(rt router) *RouterSwapper {
	rs := new(RouterSwapper)
	rs.rt = rt
	return rs
}

// ServeHTTP method for the RouterSwapper
func (rs *RouterSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.RLock()
	rt := rs.rt
	rs.mu.RUnlock()

	rt.ServeHTTP(w, r)
}
