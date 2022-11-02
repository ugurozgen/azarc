## Build stage
FROM golang:1.17.6-alpine3.15 as builder
RUN apk add --update git
ENV CGO_ENABLED 0
WORKDIR /app
COPY . .
RUN go build -a -v -o azarc .

## Run stage
FROM gcr.io/distroless/static
COPY --from=builder /app/azarc /app/azarc
CMD ["/app/azarc"]