# PROJECT IS ARCHIVED

Archived in favor of [cristalhq/appx](https://github.com/cristalhq/appx)

# sigreload

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Very small Go package to listen to `SIGHUP` signals.

## Features

* Simple API.
* Easy to integrate.
* Clean and tested code.
* Dependency-free.

## Install

Go version 1.17+

```
go get github.com/cristalhq/sigreload
```

## Example

```go
ctx := context.Background() // or something better like: signal.NotifyContext

sigreload.Do(ctx, func(ctx context.Context) {
	println("hey, time to reload config")
})
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/sigreload/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/sigreload/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/sigreload
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/sigreload
[reportcard-img]: https://goreportcard.com/badge/cristalhq/sigreload
[reportcard-url]: https://goreportcard.com/report/cristalhq/sigreload
[coverage-img]: https://codecov.io/gh/cristalhq/sigreload/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/sigreload
[version-img]: https://img.shields.io/github/v/release/cristalhq/sigreload
[version-url]: https://github.com/cristalhq/sigreload/releases
