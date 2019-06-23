package routerswapper

import (
	"net/http"
	"sync"
)

// router interface is satisfied by any type that exports ServeHTTP
type router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// RouterHandler satisfies the Router interface
type RouterHandler struct {
	router
}

// RouterSwapper is our only type
type RouterSwapper struct {
	mu sync.RWMutex
	rt *RouterHandler
}

// Swap replaces the current RouterHandler
func (rs *RouterSwapper) Swap(rt *RouterHandler) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// NewRouterSwapper creates a new RouteSwapper based on the passer RouterHandler
func NewRouterSwapper(rt *RouterHandler) (rs *RouterSwapper) {
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
