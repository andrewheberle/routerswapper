// Package routerswapper implements a simple process to swap a Golang HTTP
// router (gorilla/mux etc) during runtime. This is aimed at allowing
// route/handler changes based on a configuration change or update.
package routerswapper

import (
	"net/http"
	"sync"
)

// A Router responds to an HTTP request.
//
// Any http.Handler compatible router/mux will satisfy this interface.
type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// A RouterSwapper can be used in place of a standard mux/router by
// http.ListenAndServe() or http.ListenAndServeTLS().
//
// The RouterSwapper must be initialised with a call to 'New'.
//
// After creation, changes to the underlying mux/router must be handled
// by a call to 'Swap' to ensure the change is done in a safe manner with a
// lock.
type RouterSwapper struct {
	mu sync.RWMutex
	rt Router
}

// Swap replaces the current router with a new version ensuring a lock is
// taken so the swap is safe for concurrent use.
func (rs *RouterSwapper) Swap(rt Router) {
	rs.mu.Lock()
	rs.rt = rt
	rs.mu.Unlock()
}

// New creates a new RouterSwapper based on the provided router which can
// then be used where ServeHTTP would be used, such as http.ListenAndServe()
func New(rt Router) *RouterSwapper {
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
