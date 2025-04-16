# Dockerfile
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o arcp ./agent/main.go

FROM gcr.io/distroless/base
COPY --from=builder /app/arcp /arcp
ENTRYPOINT ["/arcp"]

