# Azarc Assignment
## How to build
```bash
docker build -t azarc:$version .
```

## How to run
```bash
docker run --rm -it -v $(pwd):/app azarc:$version
```