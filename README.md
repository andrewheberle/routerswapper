# andrewheberle/routerswapper

Package `andrewheberle/routerswapper` implements a simple process to swap a Golang HTTP router (gorilla/mux etc) during runtime. This is aimed at allowing route changes based on a configuration change or update.

Any HTTP router that satisfies the `ServeHTTP` method is supported.

As the `RouterSwapper` type implements the `ServeHTTP` method, it can be used as part of `http.ListenAndServe` etc.

## Install

```sh
go get -u gitlab.com/andrewheberle/routerswapper
```

## Examples

Using `gorilla/mux`:

```go
r := mux.NewRouter()
r.HandleFunc("/", firstHomeHandler)

rs := routerswapper.New(r)

// Contrived function to Swap the Router after 60 seconds
go func() {
    time.Sleep(60 * time.Second)
    r := mux.NewRouter()
    r.HandleFunc("/", secondHomeHandler)
    rs.Swap(r)
}()

if err := http.ListenAndServe(":8080", rs); err != nil {
	log.Fatalf("http error: %s", err)
}
```

In this example a new `mux.Router` with different a `HandleFunc` is swapped after 60 seconds.

Using `julienschmidt/httprouter`:

```go
r := httprouter.New()
r.GET("/", firstHomeHandler)

rs := routerswapper.New(r)

// Contrived function to Swap the Router after 60 seconds
go func() {
    time.Sleep(60 * time.Second)
    r := httprouter.New()
    r.GET("/", secondHomeHandler)
    rs.Swap(r)
}()

if err := http.ListenAndServe(":8080", rs); err != nil {
	log.Fatalf("http error: %s", err)
}
```