package routeswapper

import (
	"sync"
	"net/http"
)

// router interface is satisfied by any type that exports ServeHTTP
type router interface {
    ServeHTTP(w, r)
}

// RouterSwapper is our only type
type RouterSwapper struct {
	mu sync.RWMutex
	rt *Router
}

// Swap replaces the current router
func (rs *RouterSwapper) Swap(rt *Router) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// NewRouteSwapper creates a new RouteSwapper based on the passer Router
func NewRouteSwapper(rt *Router) (rs *RouterSwapper) {
	rs.rt = rt
	return rs
}

// ServeHTTP satisfies the Router interface
func (rs *RouterSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.RLock()
	rt := rs.rt
	rs.mu.RUnlock()

	rt.ServeHTTP(w, r)
}