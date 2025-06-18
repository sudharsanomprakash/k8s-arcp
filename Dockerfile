# Dockerfile
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o arcp ./agent/main.go

FROM gcr.io/distroless/base
COPY --from=builder /app/arcp /arcp
ENTRYPOINT ["/arcp"]

FROM python:3.10-bullseye

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential \
        libglib2.0-0 \
        libsm6 \
        libxext6 \
        libxrender-dev \
    && rm -rf /var/lib/apt/lists/*

RUN pip install --no-cache-dir \
    torch==2.2.1 \
    flask==3.1.1 \
    "numpy<2"

WORKDIR /app
