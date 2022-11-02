## Build stage
FROM golang:1.19-alpine3.15 as build

WORKDIR /app

ENV GOPROXY "https://proxy.golang.org,direct"
ENV GOPRIVATE "github.com/ugurozgen/azarc"
ENV CGO_ENABLED=0

COPY . .
RUN go build -a -v -o azarc 

## Run stage
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /app/azarc /azarc
USER nonroot:nonroot
ENTRYPOINT ["/azarc"]