# Azarc Assignment

### Environment variables 
`OMDB_API_KEY="b020b1bb"`

### Arguments 
```
    - filePath - absolute path to the inflated `title.basics.tsv.gz` file
    - titleType - filter on `titleType` column
    - primaryTitle - filter on `primaryTitle` column
    - originalTitle - filter on `originalTitle` column
    - genre - filter on `genre` column
    - startYear - filter on `startYear` column
    - endYear - filter on `endYear` column
    - runtimeMinutes - filter on `runtimeMinutes` column
    - genres - filter on `genres` column
    - maxApiRequests - maximum number of requests to be made to [omdbapi](https://www.omdbapi.com/)
    - maxRunTime - maximum run time of the application. Format is a `time.Duration` string see [here](https://godoc.org/time#ParseDuration)
    - maxRequests - maximum number of requests to send to [omdbapi](https://www.omdbapi.com/)
    - plotFilter - regex pattern to apply to the plot of a film retrieved from [omdbapi](https://www.omdbapi.com/)
```

## How to run tests
```bash
go test -v -coverprofile ./...
```

## How to run benchmarks
```bash
go test -benchmem -bench=.
```

## How to run cli
```bash
OMDB_API_KEY="b020b1bb" go run . -filePath ./title.basics_test.tsv -maxRunTime 10s -primaryTitle car 
```

## How to profile
```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go tool pprof -http :8080 cpu.prof
```
## Docker build
```bash
docker build -t azarc:$version .
```

## Docker run
```bash
OMDB_API_KEY="b020b1bb" docker run \
    --rm -it -v $(pwd):/app azarc:$version \
    -filePath /app/title.basics_test.tsv \
    -maxRunTime 10s \ 
    -primaryTitle car 
```