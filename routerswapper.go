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

// Swap replaces the current router
func (rs *RouterSwapper) Swap(rt router) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// New creates a new RouteSwapper based on the provided router
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
