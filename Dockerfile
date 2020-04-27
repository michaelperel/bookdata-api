FROM golang:1.14 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o app

FROM scratch
COPY --from=builder /app/app .
ENTRYPOINT ["/app"]
