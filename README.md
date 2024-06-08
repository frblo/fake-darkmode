# fake-darkmode
Very simple http server to return a very simple bool. It is meant to mock [https://github.com/datasektionen/darkmode](https://github.com/datasektionen/darkmode).

## Running locally

The server runs on port `8090`. To run it, simple run

```sh
go run . [bool]
```

where `[bool]` is either ø, `false` or `true`. In case of no argument or a bad spelling it will default to `false`.
