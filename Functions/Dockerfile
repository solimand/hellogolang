FROM golang:1.19.3-alpine AS build
WORKDIR /src/
COPY Fibonacci-funcClosure.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/fibo
FROM scratch
COPY --from=build /bin/fibo /fibo
ENTRYPOINT ["/fibo"]
