package routerswapper

import (
	"net/http"
	"sync"
)

// Router interface is satisfied by any type that exports ServeHTTP
type Router interface {
	http.Handler
}

// RouterSwapper is our only type
type RouterSwapper struct {
	mu sync.RWMutex
	rt *RouterHandler
}

// Swap replaces the current Router
func (rs *RouterSwapper) Swap(rt *Router) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// NewRouterSwapper creates a new RouteSwapper based on the passed Router
func NewRouterSwapper(rt *Router) (rs *RouterSwapper) {
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
