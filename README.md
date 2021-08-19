# Golang Libraries [![License](https://img.shields.io/badge/license-BSD%202--Clause-green.svg)](https://opensource.org/licenses/BSD-2-Clause) [![Build Status](https://github.com/rbee3u/golib/actions/workflows/build.yml/badge.svg)](https://github.com/rbee3u/golib/actions/workflows/build.yml?query=branch%3Amain)

`golib` is a collection of small, reusable Go packages.
It is a toolbox, not a framework.

- probabilistic data structures
- in-memory memoization and caching helpers
- channel utilities
- service runner helpers
- an experimental STL-inspired package set

## Packages

### `bloomfilter`

Space-efficient membership checks.

- create a filter from estimated item count and false-positive rate
- add items and test membership
- customizable hash function support

### `memo`

A concurrent in-memory key/value store with lazy loading and expiration.

- generic API
- concurrent `Get`, `Set`, and `Del`
- optional loader function for cache-miss population
- per-memo and per-call expiration settings
- duplicate concurrent loads for the same key are collapsed

### `ibch`

An "infinite-buffer" channel built from a send channel and a receive channel.

### `runner`

A thin wrapper around [`suture`](https://github.com/thejerf/suture) for supervised background services.

- simpler entry point for common supervision use cases
- signal handling for graceful shutdown
- access to restart and backoff behavior through options

### `stl`

An experimental attempt to build C++ STL-like collections and algorithms in Go.

Current contents include:

- sequence containers such as vector-like, singly linked, and doubly linked lists
- associative containers such as set, multiset, dict, and multidict
- container adaptors such as stack, queue, and deque
- a small set of generic algorithms and iterator-related constraints

## Status

This repository is usable as a utility collection, but `stl` is paused.

Go generics are helpful, but still too limited to make an STL-style library feel as expressive or elegant as C++. Many designs that are natural in C++ become awkward in Go.

Treat `stl` as:

- an experiment
- a partial implementation
- a half-finished prototype rather than a polished general-purpose library

It is useful mainly as an experiment or reference, not as a complete STL replacement.

## Repository layout

```text
.
├── bloomfilter/
├── ibch/
├── memo/
├── runner/
└── stl/
```

Packages are meant to be imported independently.

## Getting started

Import only the package you need:

```go
import "github.com/rbee3u/golib/memo"
```

Run the test suite:

```bash
go test ./...
```

Examples live in each package's `example_test.go`.

## Expectations

- APIs outside `stl` are the most usable parts of the repo
- APIs inside `stl` may feel uneven and may change without much ceremony
- this repository is best treated as a toolbox

## License

BSD 2-Clause. See [LICENSE](LICENSE).
