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

Test | Total Number of Iterations | Execution Time per operation
---- | -------------------------- | ----------------------------
BenchmarkCache-2 | 3000000 | 554 ns/op
BenchmarkRedis-2 | 10000 | 171351 ns/op