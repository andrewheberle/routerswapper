# andrewheberle/routerswapper

[![GoDoc](https://godoc.org/gitlab.com/andrewheberle/routerswapper?status.svg)](http://godoc.org/gitlab.com/andrewheberle/routerswapper)

Package `andrewheberle/routerswapper` implements a simple process to swap a Golang HTTP router (`net/http`, `gorilla/mux` etc) during runtime. This is aimed at allowing route changes based on a configuration change or update.

Any HTTP router that satisfies the `ServeHTTP` method is supported.

As the `RouterSwapper` type implements the `ServeHTTP` method, it can be used as part of `http.ListenAndServe` etc.

## Install

```sh
go get -u gitlab.com/andrewheberle/routerswapper
```

## Examples

Using `net/http`:

```go
r := http.NewServeMux()
r.HandleFunc("/", firstHomeHandler)

rs := routerswapper.New(r)

// Contrived function to Swap the Router after 60 seconds
go func() {
    time.Sleep(60 * time.Second)
    r := http.NewServeMux()
    r.HandleFunc("/", secondHomeHandler)
    rs.Swap(r)
}()

if err := http.ListenAndServe(":8080", rs); err != nil {
	log.Fatalf("http error: %s", err)
}
```

Using `github.com/gorilla/mux`:

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

Using `github.com/julienschmidt/httprouter`:

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

## Credits

The code for this package came about from this `gorilla/mux` issue: [deleting routes](https://github.com/gorilla/mux/issues/82)

## License

This project is licensed under the terms of the MIT license.
