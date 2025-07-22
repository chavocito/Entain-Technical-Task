FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

USER 1000

EXPOSE 8080

CMD ["./main"]
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s CMD curl -f http://localhost:8080/health || exit 1