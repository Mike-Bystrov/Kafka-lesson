FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

ARG APP=producer
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/${APP}

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/app /app

ENTRYPOINT ["/app"]
