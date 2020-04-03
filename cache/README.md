# High Performant Cache

High performant cache implementation

### Prerequisite
* [Go](https://golang.org/doc/install) -- Development environment
* [dep](https://golang.github.io/dep/docs/installation.html) -- Dependency management
* [Docker](https://www.docker.com/products/docker-desktop)
* [Docker Compose](https://docs.docker.com/compose/install/)

## How to run
```sh
git clone https://github.com/melvinodsa/go-experiments.git
cd cache
dep ensure
docker-compose up
```

## Benchmark Results

Test | Total Number of Iterations | Execution Time per operation | Memory allocations
---- | -------------------------- | ---------------------------- | ------------------
BenchmarkCache-4 | 1000000 | 2782 ns/op | 2 allocs/op
BenchmarkMutexCache-4 | 3336343 | 2952 ns/op | 1 allocs/op
BenchmarkRedis-4 | 575284 | 1741 ns/op | 4 allocs/op