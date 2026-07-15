FROM golang:1.22-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /taskflow-api ./cmd/api

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /taskflow-api ./taskflow-api
ENV PORT=8080
ENV DATABASE_DSN=/app/data/taskflow.db
EXPOSE 8080
CMD ["./taskflow-api"]
