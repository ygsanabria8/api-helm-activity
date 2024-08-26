# Stage 1: Build go app
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-golang

# Stage 2: Build final image
FROM gcr.io/distroless/base-debian11

COPY --from=builder /api-golang /api-golang

EXPOSE 11000
ENTRYPOINT ["/api-golang"]